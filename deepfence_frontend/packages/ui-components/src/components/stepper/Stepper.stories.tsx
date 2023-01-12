import { ComponentMeta, ComponentStory } from '@storybook/react';

import { Step, Stepper } from '@/components/stepper/Stepper';

export default {
  title: 'Components/Stepper',
  component: Stepper,
} as ComponentMeta<typeof Stepper>;

const Template: ComponentStory<typeof Stepper> = (args) => (
  <Stepper>
    <Step indicator="1" title="Ordered">
      <div>A Laptop</div>
    </Step>
    <Step indicator="2" title="Shipped">
      <div className="dark:text-gray-400 text-sm">
        Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem
        Ipsum has been the industrys standard dummy text ever since the 1500s, when an
        unknown printer took a galley of type and scrambled it to make a type specimen
        book. It has survived not only five centuries, but also the leap into electronic
        typesetting, remaining essentially unchanged. It was popularised in the 1960s with
        the release of Letraset sheets containing Lorem Ipsum passages, and more recently
        with desktop publishing software like Aldus PageMaker including versions of Lorem
        Ipsum.
      </div>
    </Step>
    <Step indicator="3" title="Delivered">
      <div className="dark:text-gray-400 text-lg">Successfully delivered</div>
    </Step>
  </Stepper>
);
export const Default = Template.bind({});

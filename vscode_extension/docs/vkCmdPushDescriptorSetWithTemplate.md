# vkCmdPushDescriptorSetWithTemplate(3) Manual Page

## Name

vkCmdPushDescriptorSetWithTemplate - Pushes descriptor updates into a command buffer using a descriptor update template



## [](#_c_specification)C Specification

It is also possible to use a descriptor update template to specify the push descriptors to update. To do so, call:

```c++
// Provided by VK_VERSION_1_4
void vkCmdPushDescriptorSetWithTemplate(
    VkCommandBuffer                             commandBuffer,
    VkDescriptorUpdateTemplate                  descriptorUpdateTemplate,
    VkPipelineLayout                            layout,
    uint32_t                                    set,
    const void*                                 pData);
```

or the equivalent command

```c++
// Provided by VK_KHR_descriptor_update_template with VK_KHR_push_descriptor, VK_KHR_push_descriptor with VK_VERSION_1_1 or VK_KHR_descriptor_update_template
void vkCmdPushDescriptorSetWithTemplateKHR(
    VkCommandBuffer                             commandBuffer,
    VkDescriptorUpdateTemplate                  descriptorUpdateTemplate,
    VkPipelineLayout                            layout,
    uint32_t                                    set,
    const void*                                 pData);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer that the descriptors will be recorded in.
- `descriptorUpdateTemplate` is a descriptor update template defining how to interpret the descriptor information in `pData`.
- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) object used to program the bindings. It **must** be compatible with the layout used to create the `descriptorUpdateTemplate` handle.
- `set` is the set number of the descriptor set in the pipeline layout that will be updated. This **must** be the same number used to create the `descriptorUpdateTemplate` handle.
- `pData` is a pointer to memory containing descriptors for the templated update.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCmdPushDescriptorSetWithTemplate-commandBuffer-00366)VUID-vkCmdPushDescriptorSetWithTemplate-commandBuffer-00366  
  The `pipelineBindPoint` specified during the creation of the descriptor update template **must** be supported by the `commandBuffer`’s parent `VkCommandPool`’s queue family
- [](#VUID-vkCmdPushDescriptorSetWithTemplate-pData-01686)VUID-vkCmdPushDescriptorSetWithTemplate-pData-01686  
  `pData` **must** be a valid pointer to a memory containing one or more valid instances of [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorImageInfo.html), [VkDescriptorBufferInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferInfo.html), or [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) in a layout defined by `descriptorUpdateTemplate` when it was created with [vkCreateDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDescriptorUpdateTemplate.html)
- [](#VUID-vkCmdPushDescriptorSetWithTemplate-layout-07993)VUID-vkCmdPushDescriptorSetWithTemplate-layout-07993  
  `layout` **must** be compatible with the layout used to create `descriptorUpdateTemplate`
- [](#VUID-vkCmdPushDescriptorSetWithTemplate-descriptorUpdateTemplate-07994)VUID-vkCmdPushDescriptorSetWithTemplate-descriptorUpdateTemplate-07994  
  `descriptorUpdateTemplate` **must** have been created with a `templateType` of `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS`
- [](#VUID-vkCmdPushDescriptorSetWithTemplate-set-07995)VUID-vkCmdPushDescriptorSetWithTemplate-set-07995  
  `set` **must** be the same value used to create `descriptorUpdateTemplate`
- [](#VUID-vkCmdPushDescriptorSetWithTemplate-set-07304)VUID-vkCmdPushDescriptorSetWithTemplate-set-07304  
  `set` **must** be less than [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount` provided when `layout` was created
- [](#VUID-vkCmdPushDescriptorSetWithTemplate-set-07305)VUID-vkCmdPushDescriptorSetWithTemplate-set-07305  
  `set` **must** be the unique set number in the pipeline layout that uses a descriptor set layout that was created with `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT`
- [](#VUID-vkCmdPushDescriptorSetWithTemplate-None-10358)VUID-vkCmdPushDescriptorSetWithTemplate-None-10358  
  If the [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html) extension is not enabled, [`pushDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pushDescriptor) **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdPushDescriptorSetWithTemplate-commandBuffer-parameter)VUID-vkCmdPushDescriptorSetWithTemplate-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdPushDescriptorSetWithTemplate-descriptorUpdateTemplate-parameter)VUID-vkCmdPushDescriptorSetWithTemplate-descriptorUpdateTemplate-parameter  
  `descriptorUpdateTemplate` **must** be a valid [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplate.html) handle
- [](#VUID-vkCmdPushDescriptorSetWithTemplate-layout-parameter)VUID-vkCmdPushDescriptorSetWithTemplate-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-vkCmdPushDescriptorSetWithTemplate-commandBuffer-recording)VUID-vkCmdPushDescriptorSetWithTemplate-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdPushDescriptorSetWithTemplate-commandBuffer-cmdpool)VUID-vkCmdPushDescriptorSetWithTemplate-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdPushDescriptorSetWithTemplate-videocoding)VUID-vkCmdPushDescriptorSetWithTemplate-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdPushDescriptorSetWithTemplate-commonparent)VUID-vkCmdPushDescriptorSetWithTemplate-commonparent  
  Each of `commandBuffer`, `descriptorUpdateTemplate`, and `layout` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Outside

Graphics  
Compute

State

Conditional Rendering

vkCmdPushDescriptorSetWithTemplate is not affected by [conditional rendering](#drawing-conditional-rendering)

API Example

```c++
struct AppDataStructure
{
    VkDescriptorImageInfo  imageInfo;          // a single image info
    // ... some more application-related data
};

const VkDescriptorUpdateTemplateEntry descriptorUpdateTemplateEntries[] =
{
    // binding to a single image descriptor
    {
        .binding = 0,
        .dstArrayElement = 0,
        .descriptorCount = 1,
        .descriptorType = VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER,
        .offset = offsetof(AppDataStructure, imageInfo),
        .stride = 0     // not required if descriptorCount is 1
    }
};

// create a descriptor update template for push descriptor set updates
const VkDescriptorUpdateTemplateCreateInfo createInfo =
{
    .sType = VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO,
    .pNext = NULL,
    .flags = 0,
    .descriptorUpdateEntryCount = 1,
    .pDescriptorUpdateEntries = descriptorUpdateTemplateEntries,
    .templateType = VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS,
    .descriptorSetLayout = 0,   // ignored by given templateType
    .pipelineBindPoint = VK_PIPELINE_BIND_POINT_GRAPHICS,
    .pipelineLayout = myPipelineLayout,
    .set = 0,
};

VkDescriptorUpdateTemplate myDescriptorUpdateTemplate;
myResult = vkCreateDescriptorUpdateTemplate(
    myDevice,
    &createInfo,
    NULL,
    &myDescriptorUpdateTemplate);

AppDataStructure appData;
// fill appData here or cache it in your engine
vkCmdPushDescriptorSetWithTemplate(myCmdBuffer, myDescriptorUpdateTemplate, myPipelineLayout, 0,&appData);
```

## [](#_see_also)See Also

[VK\_KHR\_descriptor\_update\_template](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_descriptor_update_template.html), [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html), [VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplate.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdPushDescriptorSetWithTemplate)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
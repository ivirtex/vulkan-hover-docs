# vkCmdPushDescriptorSet(3) Manual Page

## Name

vkCmdPushDescriptorSet - Pushes descriptor updates into a command buffer



## [](#_c_specification)C Specification

In addition to allocating descriptor sets and binding them to a command buffer, an application **can** record descriptor updates into the command buffer.

To push descriptor updates into a command buffer, call:

```c++
// Provided by VK_VERSION_1_4
void vkCmdPushDescriptorSet(
    VkCommandBuffer                             commandBuffer,
    VkPipelineBindPoint                         pipelineBindPoint,
    VkPipelineLayout                            layout,
    uint32_t                                    set,
    uint32_t                                    descriptorWriteCount,
    const VkWriteDescriptorSet*                 pDescriptorWrites);
```

or the equivalent command

```c++
// Provided by VK_KHR_push_descriptor
void vkCmdPushDescriptorSetKHR(
    VkCommandBuffer                             commandBuffer,
    VkPipelineBindPoint                         pipelineBindPoint,
    VkPipelineLayout                            layout,
    uint32_t                                    set,
    uint32_t                                    descriptorWriteCount,
    const VkWriteDescriptorSet*                 pDescriptorWrites);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer that the descriptors will be recorded in.
- `pipelineBindPoint` is a [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) indicating the type of the pipeline that will use the descriptors. There is a separate set of push descriptor bindings for each pipeline type, so binding one does not disturb the others.
- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) object used to program the bindings.
- `set` is the set number of the descriptor set in the pipeline layout that will be updated.
- `descriptorWriteCount` is the number of elements in the `pDescriptorWrites` array.
- `pDescriptorWrites` is a pointer to an array of [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html) structures describing the descriptors to be updated.

## [](#_description)Description

*Push descriptors* are a small bank of descriptors whose storage is internally managed by the command buffer rather than being written into a descriptor set and later bound to a command buffer. Push descriptors allow for incremental updates of descriptors without managing the lifetime of descriptor sets.

When a command buffer begins recording, all push descriptors are undefined. Push descriptors **can** be updated incrementally and cause shaders to use the updated descriptors for subsequent [bound pipeline commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-bindpoint-commands) with the pipeline type set by `pipelineBindPoint` until the descriptor is overwritten, or else until the set is disturbed as described in [Pipeline Layout Compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-compatibility). When the set is disturbed or push descriptors with a different descriptor set layout are set, all push descriptors are undefined.

Push descriptors that are [statically used](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-staticuse) by a pipeline **must** not be undefined at the time that a drawing or dispatching command is recorded to execute using that pipeline. This includes immutable sampler descriptors, which **must** be pushed before they are accessed by a pipeline (the immutable samplers are pushed, rather than the samplers in `pDescriptorWrites`). Push descriptors that are not statically used **can** remain undefined.

Push descriptors do not use dynamic offsets. Instead, the corresponding non-dynamic descriptor types **can** be used and the `offset` member of [VkDescriptorBufferInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferInfo.html) **can** be changed each time the descriptor is written.

Each element of `pDescriptorWrites` is interpreted as in [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html), except the `dstSet` member is ignored.

To push an immutable sampler, use a [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html) with `dstBinding` and `dstArrayElement` selecting the immutable sampler’s binding. If the descriptor type is `VK_DESCRIPTOR_TYPE_SAMPLER`, the `pImageInfo` parameter is ignored and the immutable sampler is taken from the push descriptor set layout in the pipeline layout. If the descriptor type is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, the `sampler` member of the `pImageInfo` parameter is ignored and the immutable sampler is taken from the push descriptor set layout in the pipeline layout.

Valid Usage

- [](#VUID-vkCmdPushDescriptorSet-set-00364)VUID-vkCmdPushDescriptorSet-set-00364  
  `set` **must** be less than [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount` provided when `layout` was created
- [](#VUID-vkCmdPushDescriptorSet-set-00365)VUID-vkCmdPushDescriptorSet-set-00365  
  `set` **must** be the unique set number in the pipeline layout that uses a descriptor set layout that was created with `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT`
- [](#VUID-vkCmdPushDescriptorSet-pDescriptorWrites-06494)VUID-vkCmdPushDescriptorSet-pDescriptorWrites-06494  
  For each element i where `pDescriptorWrites`\[i].`descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLER`, `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, `pDescriptorWrites`\[i].`pImageInfo` **must** be a valid pointer to an array of `pDescriptorWrites`\[i].`descriptorCount` valid `VkDescriptorImageInfo` structures
- [](#VUID-vkCmdPushDescriptorSet-pipelineBindPoint-00363)VUID-vkCmdPushDescriptorSet-pipelineBindPoint-00363  
  `pipelineBindPoint` **must** be supported by the `commandBuffer`’s parent `VkCommandPool`’s queue family
- [](#VUID-vkCmdPushDescriptorSet-None-10356)VUID-vkCmdPushDescriptorSet-None-10356  
  If the [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html) extension is not enabled, [`pushDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pushDescriptor) **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdPushDescriptorSet-commandBuffer-parameter)VUID-vkCmdPushDescriptorSet-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdPushDescriptorSet-pipelineBindPoint-parameter)VUID-vkCmdPushDescriptorSet-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value
- [](#VUID-vkCmdPushDescriptorSet-layout-parameter)VUID-vkCmdPushDescriptorSet-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-vkCmdPushDescriptorSet-pDescriptorWrites-parameter)VUID-vkCmdPushDescriptorSet-pDescriptorWrites-parameter  
  `pDescriptorWrites` **must** be a valid pointer to an array of `descriptorWriteCount` valid [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html) structures
- [](#VUID-vkCmdPushDescriptorSet-commandBuffer-recording)VUID-vkCmdPushDescriptorSet-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdPushDescriptorSet-commandBuffer-cmdpool)VUID-vkCmdPushDescriptorSet-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdPushDescriptorSet-videocoding)VUID-vkCmdPushDescriptorSet-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdPushDescriptorSet-descriptorWriteCount-arraylength)VUID-vkCmdPushDescriptorSet-descriptorWriteCount-arraylength  
  `descriptorWriteCount` **must** be greater than `0`
- [](#VUID-vkCmdPushDescriptorSet-commonparent)VUID-vkCmdPushDescriptorSet-commonparent  
  Both of `commandBuffer`, and `layout` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

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

vkCmdPushDescriptorSet is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html), [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdPushDescriptorSet).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
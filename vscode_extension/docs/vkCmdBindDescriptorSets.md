# vkCmdBindDescriptorSets(3) Manual Page

## Name

vkCmdBindDescriptorSets - Binds descriptor sets to a command buffer



## [](#_c_specification)C Specification

To bind one or more descriptor sets to a command buffer, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdBindDescriptorSets(
    VkCommandBuffer                             commandBuffer,
    VkPipelineBindPoint                         pipelineBindPoint,
    VkPipelineLayout                            layout,
    uint32_t                                    firstSet,
    uint32_t                                    descriptorSetCount,
    const VkDescriptorSet*                      pDescriptorSets,
    uint32_t                                    dynamicOffsetCount,
    const uint32_t*                             pDynamicOffsets);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer that the descriptor sets will be bound to.
- `pipelineBindPoint` is a [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) indicating the type of the pipeline that will use the descriptors. There is a separate set of bind points for each pipeline type, so binding one does not disturb the others.
- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) object used to program the bindings.
- `firstSet` is the set number of the first descriptor set to be bound.
- `descriptorSetCount` is the number of elements in the `pDescriptorSets` array.
- `pDescriptorSets` is a pointer to an array of handles to [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html) objects describing the descriptor sets to bind to.
- `dynamicOffsetCount` is the number of dynamic offsets in the `pDynamicOffsets` array.
- `pDynamicOffsets` is a pointer to an array of `uint32_t` values specifying dynamic offsets.

## [](#_description)Description

`vkCmdBindDescriptorSets` binds descriptor sets `pDescriptorSets`\[0..`descriptorSetCount`-1] to set numbers \[`firstSet`..`firstSet`+`descriptorSetCount`-1] for subsequent [bound pipeline commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-bindpoint-commands) set by `pipelineBindPoint`. Any bindings that were previously applied via these sets , or calls to [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html) or [vkCmdBindDescriptorBufferEmbeddedSamplersEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorBufferEmbeddedSamplersEXT.html), are no longer valid.

Once bound, a descriptor set affects rendering of subsequent commands that interact with the given pipeline type in the command buffer until either a different set is bound to the same set number, or the set is disturbed as described in [Pipeline Layout Compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-compatibility).

A compatible descriptor set **must** be bound for all set numbers that any shaders in a pipeline access, at the time that a drawing or dispatching command is recorded to execute using that pipeline. However, if none of the shaders in a pipeline statically use any bindings with a particular set number, then no descriptor set need be bound for that set number, even if the pipeline layout includes a non-trivial descriptor set layout for that set number.

When consuming a descriptor, a descriptor is considered valid if the descriptor is not undefined as described by [descriptor set allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptor-set-initial-state). If the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is enabled, a null descriptor is also considered valid. A descriptor that was disturbed by [Pipeline Layout Compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-compatibility), or was never bound by `vkCmdBindDescriptorSets` is not considered valid. If a pipeline accesses a descriptor either statically or dynamically depending on the [VkDescriptorBindingFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBindingFlagBits.html), the consuming descriptor type in the pipeline **must** match the [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html) in [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html) for the descriptor to be considered valid. If a descriptor is a mutable descriptor, the consuming descriptor type in the pipeline **must** match the active descriptor type for the descriptor to be considered valid.

Note

Further validation may be carried out beyond validation for descriptor types, e.g. [Texel Input Validation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-input-validation).

If any of the sets being bound include dynamic uniform or storage buffers, then `pDynamicOffsets` includes one element for each array element in each dynamic descriptor type binding in each set. Values are taken from `pDynamicOffsets` in an order such that all entries for set N come before set N+1; within a set, entries are ordered by the binding numbers in the descriptor set layouts; and within a binding array, elements are in order. `dynamicOffsetCount` **must** equal the total number of dynamic descriptors in the sets being bound.

The effective offset used for dynamic uniform and storage buffer bindings is the sum of the relative offset taken from `pDynamicOffsets`, and the base address of the buffer plus base offset in the descriptor set. The range of the dynamic uniform and storage buffer bindings is the buffer range as specified in the descriptor set.

Each of the `pDescriptorSets` **must** be compatible with the pipeline layout specified by `layout`. The layout used to program the bindings **must** also be compatible with the pipeline used in subsequent [bound pipeline commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-bindpoint-commands) with that pipeline type, as defined in the [Pipeline Layout Compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-compatibility) section.

The descriptor set contents bound by a call to `vkCmdBindDescriptorSets` **may** be consumed at the following times:

- For descriptor bindings created with the `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` bit set, the contents **may** be consumed when the command buffer is submitted to a queue, or during shader execution of the resulting draws and dispatches, or any time in between. Otherwise,
- during host execution of the command, or during shader execution of the resulting draws and dispatches, or any time in between.

Thus, the contents of a descriptor set binding **must** not be altered (overwritten by an update command, or freed) between the first point in time that it **may** be consumed, and when the command completes executing on the queue.

The contents of `pDynamicOffsets` are consumed immediately during execution of `vkCmdBindDescriptorSets`. Once all pending uses have completed, it is legal to update and reuse a descriptor set.

Valid Usage

- [](#VUID-vkCmdBindDescriptorSets-pDescriptorSets-00358)VUID-vkCmdBindDescriptorSets-pDescriptorSets-00358  
  Each element of `pDescriptorSets` that is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) **must** have been allocated with a `VkDescriptorSetLayout` that matches (is the same as, or identically defined as) the `VkDescriptorSetLayout` at set *n* in `layout`, where *n* is the sum of `firstSet` and the index into `pDescriptorSets`
- [](#VUID-vkCmdBindDescriptorSets-dynamicOffsetCount-00359)VUID-vkCmdBindDescriptorSets-dynamicOffsetCount-00359  
  `dynamicOffsetCount` **must** be equal to the total number of dynamic descriptors in `pDescriptorSets`
- [](#VUID-vkCmdBindDescriptorSets-firstSet-00360)VUID-vkCmdBindDescriptorSets-firstSet-00360  
  The sum of `firstSet` and `descriptorSetCount` **must** be less than or equal to [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount` provided when `layout` was created
- [](#VUID-vkCmdBindDescriptorSets-pDynamicOffsets-01971)VUID-vkCmdBindDescriptorSets-pDynamicOffsets-01971  
  Each element of `pDynamicOffsets` which corresponds to a descriptor binding with type `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` **must** be a multiple of `VkPhysicalDeviceLimits`::`minUniformBufferOffsetAlignment`
- [](#VUID-vkCmdBindDescriptorSets-pDynamicOffsets-01972)VUID-vkCmdBindDescriptorSets-pDynamicOffsets-01972  
  Each element of `pDynamicOffsets` which corresponds to a descriptor binding with type `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` **must** be a multiple of `VkPhysicalDeviceLimits`::`minStorageBufferOffsetAlignment`
- [](#VUID-vkCmdBindDescriptorSets-pDescriptorSets-01979)VUID-vkCmdBindDescriptorSets-pDescriptorSets-01979  
  For each dynamic uniform or storage buffer binding in `pDescriptorSets`, the sum of the [effective offset](#dynamic-effective-offset) and the range of the binding **must** be less than or equal to the size of the buffer
- [](#VUID-vkCmdBindDescriptorSets-pDescriptorSets-06715)VUID-vkCmdBindDescriptorSets-pDescriptorSets-06715  
  For each dynamic uniform or storage buffer binding in `pDescriptorSets`, if the range was set with `VK_WHOLE_SIZE` then `pDynamicOffsets` which corresponds to the descriptor binding **must** be 0
- [](#VUID-vkCmdBindDescriptorSets-pDescriptorSets-04616)VUID-vkCmdBindDescriptorSets-pDescriptorSets-04616  
  Each element of `pDescriptorSets` **must** not have been allocated from a `VkDescriptorPool` with the `VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT` flag set
- [](#VUID-vkCmdBindDescriptorSets-pDescriptorSets-06563)VUID-vkCmdBindDescriptorSets-pDescriptorSets-06563  
  If the [`graphicsPipelineLibrary`](#features-graphicsPipelineLibrary) feature is not enabled, each element of `pDescriptorSets` **must** be a valid [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html)
- [](#VUID-vkCmdBindDescriptorSets-pDescriptorSets-08010)VUID-vkCmdBindDescriptorSets-pDescriptorSets-08010  
  Each element of `pDescriptorSets` **must** have been allocated with a `VkDescriptorSetLayout` which was not created with `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-vkCmdBindDescriptorSets-pDescriptorSets-09914)VUID-vkCmdBindDescriptorSets-pDescriptorSets-09914  
  If any element of `pDescriptorSets` was allocated from a descriptor pool created with a [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html) structure that had a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure specifying foreign data processing engines in its `pNext` chain, then the command pool from which `commandBuffer` was allocated **must** have been created with a [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) structure that had a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure in its `pNext` chain specifying a superset of all the foreign data processing engines specified when creating the descriptor pools from which the elements of `pDescriptorSets` were allocated
- [](#VUID-vkCmdBindDescriptorSets-pDescriptorSets-09915)VUID-vkCmdBindDescriptorSets-pDescriptorSets-09915  
  If none of the elements of `pDescriptorSets` were allocated from a descriptor pool created with a [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html) structure that had a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure specifying foreign data processing engines in its `pNext` chain, then the command pool from which `commandBuffer` was allocated **must** not have been created with a [VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolCreateInfo.html) structure that had a [VkDataGraphProcessingEngineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphProcessingEngineCreateInfoARM.html) structure in its `pNext` chain
- [](#VUID-vkCmdBindDescriptorSets-pipelineBindPoint-00361)VUID-vkCmdBindDescriptorSets-pipelineBindPoint-00361  
  `pipelineBindPoint` **must** be supported by the `commandBuffer`’s parent `VkCommandPool`’s queue family

Valid Usage (Implicit)

- [](#VUID-vkCmdBindDescriptorSets-commandBuffer-parameter)VUID-vkCmdBindDescriptorSets-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBindDescriptorSets-pipelineBindPoint-parameter)VUID-vkCmdBindDescriptorSets-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value
- [](#VUID-vkCmdBindDescriptorSets-layout-parameter)VUID-vkCmdBindDescriptorSets-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-vkCmdBindDescriptorSets-pDescriptorSets-parameter)VUID-vkCmdBindDescriptorSets-pDescriptorSets-parameter  
  `pDescriptorSets` **must** be a valid pointer to an array of `descriptorSetCount` valid or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html) handles
- [](#VUID-vkCmdBindDescriptorSets-pDynamicOffsets-parameter)VUID-vkCmdBindDescriptorSets-pDynamicOffsets-parameter  
  If `dynamicOffsetCount` is not `0`, `pDynamicOffsets` **must** be a valid pointer to an array of `dynamicOffsetCount` `uint32_t` values
- [](#VUID-vkCmdBindDescriptorSets-commandBuffer-recording)VUID-vkCmdBindDescriptorSets-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBindDescriptorSets-commandBuffer-cmdpool)VUID-vkCmdBindDescriptorSets-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, compute, or data\_graph operations
- [](#VUID-vkCmdBindDescriptorSets-videocoding)VUID-vkCmdBindDescriptorSets-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBindDescriptorSets-descriptorSetCount-arraylength)VUID-vkCmdBindDescriptorSets-descriptorSetCount-arraylength  
  `descriptorSetCount` **must** be greater than `0`
- [](#VUID-vkCmdBindDescriptorSets-commonparent)VUID-vkCmdBindDescriptorSets-commonparent  
  Each of `commandBuffer`, `layout`, and the elements of `pDescriptorSets` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

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
Data\_Graph

State

Conditional Rendering

vkCmdBindDescriptorSets is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html), [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBindDescriptorSets)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
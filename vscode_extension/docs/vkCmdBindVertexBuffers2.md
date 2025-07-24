# vkCmdBindVertexBuffers2(3) Manual Page

## Name

vkCmdBindVertexBuffers2 - Bind vertex buffers to a command buffer and dynamically set strides



## [](#_c_specification)C Specification

Alternatively, to bind vertex buffers, along with their sizes and strides, to a command buffer for use in subsequent drawing commands, call:

```c++
// Provided by VK_VERSION_1_3
void vkCmdBindVertexBuffers2(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstBinding,
    uint32_t                                    bindingCount,
    const VkBuffer*                             pBuffers,
    const VkDeviceSize*                         pOffsets,
    const VkDeviceSize*                         pSizes,
    const VkDeviceSize*                         pStrides);
```

or the equivalent command

```c++
// Provided by VK_EXT_extended_dynamic_state, VK_EXT_shader_object
void vkCmdBindVertexBuffers2EXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstBinding,
    uint32_t                                    bindingCount,
    const VkBuffer*                             pBuffers,
    const VkDeviceSize*                         pOffsets,
    const VkDeviceSize*                         pSizes,
    const VkDeviceSize*                         pStrides);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command is recorded.
- `firstBinding` is the index of the first vertex input binding whose state is updated by the command.
- `bindingCount` is the number of vertex input bindings whose state is updated by the command.
- `pBuffers` is a pointer to an array of buffer handles.
- `pOffsets` is a pointer to an array of buffer offsets.
- `pSizes` is `NULL` or a pointer to an array of the size in bytes of vertex data bound from `pBuffers`.
- `pStrides` is `NULL` or a pointer to an array of buffer strides.

## [](#_description)Description

The values taken from elements i of `pBuffers` and `pOffsets` replace the current state for the vertex input binding `firstBinding` + i, for i in \[0, `bindingCount`). The vertex input binding is updated to start at the offset indicated by `pOffsets`\[i] from the start of the buffer `pBuffers`\[i]. If `pSizes` is not `NULL` then `pSizes`\[i] specifies the bound size of the vertex buffer starting from the corresponding elements of `pBuffers`\[i] plus `pOffsets`\[i]. If `pSizes`\[i] is `VK_WHOLE_SIZE` then the bound size is from `pBuffers`\[i] plus `pOffsets`\[i] to the end of the buffer `pBuffers`\[i]. All vertex input attributes that use each of these bindings will use these updated addresses in their address calculations for subsequent drawing commands. If the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is enabled, elements of `pBuffers` **can** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), and **can** be used by the vertex shader. If a vertex input attribute is bound to a vertex input binding that is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the values taken from memory are considered to be zero, and missing G, B, or A components are [filled with (0,0,1)](#fxvertex-input-extraction).

This command also [dynamically sets](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the byte strides between consecutive elements within buffer `pBuffers`\[i] to the corresponding `pStrides`\[i] value when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, strides are specified by the [VkVertexInputBindingDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription.html)::`stride` values used to create the currently active pipeline.

If drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects) or if the bound pipeline state object was also created with the `VK_DYNAMIC_STATE_VERTEX_INPUT_EXT` dynamic state enabled then [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetVertexInputEXT.html) **can** be used instead of `vkCmdBindVertexBuffers2` to set the stride.

Note

Unlike the static state to set the same, `pStrides` must be between 0 and the maximum extent of the attributes in the binding. [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetVertexInputEXT.html) does not have this restriction so can be used if other stride values are desired.

Valid Usage

- [](#VUID-vkCmdBindVertexBuffers2-None-08971)VUID-vkCmdBindVertexBuffers2-None-08971  
  At least one of the following **must** be true:
  
  - the [`extendedDynamicState`](#features-extendedDynamicState) feature is enabled
  - the [`shaderObject`](#features-shaderObject) feature is enabled
  - the value of [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion` used to create the [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) parent of `commandBuffer` is greater than or equal to Version 1.3
- [](#VUID-vkCmdBindVertexBuffers2-firstBinding-03355)VUID-vkCmdBindVertexBuffers2-firstBinding-03355  
  `firstBinding` **must** be less than `VkPhysicalDeviceLimits`::`maxVertexInputBindings`
- [](#VUID-vkCmdBindVertexBuffers2-firstBinding-03356)VUID-vkCmdBindVertexBuffers2-firstBinding-03356  
  The sum of `firstBinding` and `bindingCount` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxVertexInputBindings`
- [](#VUID-vkCmdBindVertexBuffers2-pOffsets-03357)VUID-vkCmdBindVertexBuffers2-pOffsets-03357  
  If `pSizes` is not `NULL`, all elements of `pOffsets` **must** be less than the size of the corresponding element in `pBuffers`
- [](#VUID-vkCmdBindVertexBuffers2-pSizes-03358)VUID-vkCmdBindVertexBuffers2-pSizes-03358  
  If `pSizes` is not `NULL`, all elements of `pOffsets` plus `pSizes` , where `pSizes` is not `VK_WHOLE_SIZE`, **must** be less than or equal to the size of the corresponding element in `pBuffers`
- [](#VUID-vkCmdBindVertexBuffers2-pBuffers-03359)VUID-vkCmdBindVertexBuffers2-pBuffers-03359  
  All elements of `pBuffers` **must** have been created with the `VK_BUFFER_USAGE_VERTEX_BUFFER_BIT` flag
- [](#VUID-vkCmdBindVertexBuffers2-pBuffers-03360)VUID-vkCmdBindVertexBuffers2-pBuffers-03360  
  Each element of `pBuffers` that is non-sparse **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdBindVertexBuffers2-pBuffers-04111)VUID-vkCmdBindVertexBuffers2-pBuffers-04111  
  If the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, all elements of `pBuffers` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCmdBindVertexBuffers2-pBuffers-04112)VUID-vkCmdBindVertexBuffers2-pBuffers-04112  
  If an element of `pBuffers` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), then the corresponding element of `pOffsets` **must** be zero
- [](#VUID-vkCmdBindVertexBuffers2-pStrides-03362)VUID-vkCmdBindVertexBuffers2-pStrides-03362  
  If `pStrides` is not `NULL` each element of `pStrides` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxVertexInputBindingStride`
- [](#VUID-vkCmdBindVertexBuffers2-pStrides-06209)VUID-vkCmdBindVertexBuffers2-pStrides-06209  
  If `pStrides` is not `NULL` each element of `pStrides` **must** be either 0 or greater than or equal to the maximum extent of all vertex input attributes fetched from the corresponding binding, where the extent is calculated as the [VkVertexInputAttributeDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputAttributeDescription.html)::`offset` plus [VkVertexInputAttributeDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputAttributeDescription.html)::`format` size

Valid Usage (Implicit)

- [](#VUID-vkCmdBindVertexBuffers2-commandBuffer-parameter)VUID-vkCmdBindVertexBuffers2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBindVertexBuffers2-pBuffers-parameter)VUID-vkCmdBindVertexBuffers2-pBuffers-parameter  
  `pBuffers` **must** be a valid pointer to an array of `bindingCount` valid or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handles
- [](#VUID-vkCmdBindVertexBuffers2-pOffsets-parameter)VUID-vkCmdBindVertexBuffers2-pOffsets-parameter  
  `pOffsets` **must** be a valid pointer to an array of `bindingCount` [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) values
- [](#VUID-vkCmdBindVertexBuffers2-pSizes-parameter)VUID-vkCmdBindVertexBuffers2-pSizes-parameter  
  If `pSizes` is not `NULL`, `pSizes` **must** be a valid pointer to an array of `bindingCount` [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) values
- [](#VUID-vkCmdBindVertexBuffers2-pStrides-parameter)VUID-vkCmdBindVertexBuffers2-pStrides-parameter  
  If `pStrides` is not `NULL`, `pStrides` **must** be a valid pointer to an array of `bindingCount` [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) values
- [](#VUID-vkCmdBindVertexBuffers2-commandBuffer-recording)VUID-vkCmdBindVertexBuffers2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBindVertexBuffers2-commandBuffer-cmdpool)VUID-vkCmdBindVertexBuffers2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdBindVertexBuffers2-videocoding)VUID-vkCmdBindVertexBuffers2-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBindVertexBuffers2-bindingCount-arraylength)VUID-vkCmdBindVertexBuffers2-bindingCount-arraylength  
  If any of `pSizes`, or `pStrides` are not `NULL`, `bindingCount` **must** be greater than `0`
- [](#VUID-vkCmdBindVertexBuffers2-commonparent)VUID-vkCmdBindVertexBuffers2-commonparent  
  Both of `commandBuffer`, and the elements of `pBuffers` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

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

State

Conditional Rendering

vkCmdBindVertexBuffers2 is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBindVertexBuffers2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
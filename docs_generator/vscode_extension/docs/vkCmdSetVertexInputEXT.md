# vkCmdSetVertexInputEXT(3) Manual Page

## Name

vkCmdSetVertexInputEXT - Set the vertex input state dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the vertex input attribute and vertex input binding descriptions, call:

```c++
// Provided by VK_EXT_shader_object, VK_EXT_vertex_input_dynamic_state
void vkCmdSetVertexInputEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    vertexBindingDescriptionCount,
    const VkVertexInputBindingDescription2EXT*  pVertexBindingDescriptions,
    uint32_t                                    vertexAttributeDescriptionCount,
    const VkVertexInputAttributeDescription2EXT* pVertexAttributeDescriptions);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `vertexBindingDescriptionCount` is the number of vertex binding descriptions provided in `pVertexBindingDescriptions`.
- `pVertexBindingDescriptions` is a pointer to an array of [VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription2EXT.html) structures.
- `vertexAttributeDescriptionCount` is the number of vertex attribute descriptions provided in `pVertexAttributeDescriptions`.
- `pVertexAttributeDescriptions` is a pointer to an array of [VkVertexInputAttributeDescription2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputAttributeDescription2EXT.html) structures.

## [](#_description)Description

This command sets the vertex input attribute and vertex input binding descriptions state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_VERTEX_INPUT_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html)::`pVertexInputState` values used to create the currently active pipeline.

If drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or if the bound pipeline state object was also created with the `VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE` dynamic state enabled, then [vkCmdBindVertexBuffers2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindVertexBuffers2.html) can be used instead of `vkCmdSetVertexInputEXT` to dynamically set the stride.

The vertex attribute description for any location in the range [0,`VkPhysicalDeviceLimits`::`maxVertexInputAttributes`) not specified in the `pVertexAttributeDescriptions` array becomes undefined.

Valid Usage

- [](#VUID-vkCmdSetVertexInputEXT-None-08546)VUID-vkCmdSetVertexInputEXT-None-08546  
  Either the [`vertexInputDynamicState`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-vertexInputDynamicState) feature or the [`shaderObject`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderObject) feature or both **must** be enabled
- [](#VUID-vkCmdSetVertexInputEXT-vertexBindingDescriptionCount-04791)VUID-vkCmdSetVertexInputEXT-vertexBindingDescriptionCount-04791  
  `vertexBindingDescriptionCount` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxVertexInputBindings`
- [](#VUID-vkCmdSetVertexInputEXT-vertexAttributeDescriptionCount-04792)VUID-vkCmdSetVertexInputEXT-vertexAttributeDescriptionCount-04792  
  `vertexAttributeDescriptionCount` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxVertexInputAttributes`
- [](#VUID-vkCmdSetVertexInputEXT-binding-04793)VUID-vkCmdSetVertexInputEXT-binding-04793  
  For every `binding` specified by each element of `pVertexAttributeDescriptions`, a [VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription2EXT.html) **must** exist in `pVertexBindingDescriptions` with the same value of `binding`
- [](#VUID-vkCmdSetVertexInputEXT-pVertexBindingDescriptions-04794)VUID-vkCmdSetVertexInputEXT-pVertexBindingDescriptions-04794  
  All elements of `pVertexBindingDescriptions` **must** describe distinct binding numbers
- [](#VUID-vkCmdSetVertexInputEXT-pVertexAttributeDescriptions-04795)VUID-vkCmdSetVertexInputEXT-pVertexAttributeDescriptions-04795  
  All elements of `pVertexAttributeDescriptions` **must** describe distinct attribute locations

Valid Usage (Implicit)

- [](#VUID-vkCmdSetVertexInputEXT-commandBuffer-parameter)VUID-vkCmdSetVertexInputEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetVertexInputEXT-pVertexBindingDescriptions-parameter)VUID-vkCmdSetVertexInputEXT-pVertexBindingDescriptions-parameter  
  If `vertexBindingDescriptionCount` is not `0`, `pVertexBindingDescriptions` **must** be a valid pointer to an array of `vertexBindingDescriptionCount` valid [VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription2EXT.html) structures
- [](#VUID-vkCmdSetVertexInputEXT-pVertexAttributeDescriptions-parameter)VUID-vkCmdSetVertexInputEXT-pVertexAttributeDescriptions-parameter  
  If `vertexAttributeDescriptionCount` is not `0`, `pVertexAttributeDescriptions` **must** be a valid pointer to an array of `vertexAttributeDescriptionCount` valid [VkVertexInputAttributeDescription2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputAttributeDescription2EXT.html) structures
- [](#VUID-vkCmdSetVertexInputEXT-commandBuffer-recording)VUID-vkCmdSetVertexInputEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetVertexInputEXT-commandBuffer-cmdpool)VUID-vkCmdSetVertexInputEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetVertexInputEXT-videocoding)VUID-vkCmdSetVertexInputEXT-videocoding  
  This command **must** only be called outside of a video coding scope

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

## [](#_see_also)See Also

[VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_EXT\_vertex\_input\_dynamic\_state](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_vertex_input_dynamic_state.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkVertexInputAttributeDescription2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputAttributeDescription2EXT.html), [VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription2EXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetVertexInputEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# vkCmdSetTessellationDomainOriginEXT(3) Manual Page

## Name

vkCmdSetTessellationDomainOriginEXT - Specify the origin of the tessellation domain space dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the origin of the tessellation domain space, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3 with VK_KHR_maintenance2 or VK_VERSION_1_1, VK_EXT_shader_object
void vkCmdSetTessellationDomainOriginEXT(
    VkCommandBuffer                             commandBuffer,
    VkTessellationDomainOrigin                  domainOrigin);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `domainOrigin` specifies the origin of the tessellation domain space.

## [](#_description)Description

This command sets the origin of the tessellation domain space for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_TESSELLATION_DOMAIN_ORIGIN_EXT` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineTessellationDomainOriginStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineTessellationDomainOriginStateCreateInfo.html)::`domainOrigin` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetTessellationDomainOriginEXT-None-09423)VUID-vkCmdSetTessellationDomainOriginEXT-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3TessellationDomainOrigin`](#features-extendedDynamicState3TessellationDomainOrigin) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetTessellationDomainOriginEXT-commandBuffer-parameter)VUID-vkCmdSetTessellationDomainOriginEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetTessellationDomainOriginEXT-domainOrigin-parameter)VUID-vkCmdSetTessellationDomainOriginEXT-domainOrigin-parameter  
  `domainOrigin` **must** be a valid [VkTessellationDomainOrigin](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTessellationDomainOrigin.html) value
- [](#VUID-vkCmdSetTessellationDomainOriginEXT-commandBuffer-recording)VUID-vkCmdSetTessellationDomainOriginEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetTessellationDomainOriginEXT-commandBuffer-cmdpool)VUID-vkCmdSetTessellationDomainOriginEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetTessellationDomainOriginEXT-videocoding)VUID-vkCmdSetTessellationDomainOriginEXT-videocoding  
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

Conditional Rendering

vkCmdSetTessellationDomainOriginEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_KHR\_maintenance2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance2.html), [VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkTessellationDomainOrigin](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTessellationDomainOrigin.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetTessellationDomainOriginEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
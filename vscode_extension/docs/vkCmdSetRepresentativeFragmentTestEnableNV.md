# vkCmdSetRepresentativeFragmentTestEnableNV(3) Manual Page

## Name

vkCmdSetRepresentativeFragmentTestEnableNV - Specify the representative fragment test enable dynamically for a command buffer



## [](#_c_specification)C Specification

To [dynamically set](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-dynamic-state) the `representativeFragmentTestEnable` state, call:

```c++
// Provided by VK_EXT_extended_dynamic_state3 with VK_NV_representative_fragment_test, VK_EXT_shader_object with VK_NV_representative_fragment_test
void vkCmdSetRepresentativeFragmentTestEnableNV(
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    representativeFragmentTestEnable);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `representativeFragmentTestEnable` specifies the `representativeFragmentTestEnable` state.

## [](#_description)Description

This command sets the `representativeFragmentTestEnable` state for subsequent drawing commands when drawing using [shader objects](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-objects), or when the graphics pipeline is created with `VK_DYNAMIC_STATE_REPRESENTATIVE_FRAGMENT_TEST_ENABLE_NV` set in [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`. Otherwise, this state is specified by the [VkPipelineRepresentativeFragmentTestStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRepresentativeFragmentTestStateCreateInfoNV.html)::`representativeFragmentTestEnable` value used to create the currently active pipeline.

Valid Usage

- [](#VUID-vkCmdSetRepresentativeFragmentTestEnableNV-None-09423)VUID-vkCmdSetRepresentativeFragmentTestEnableNV-None-09423  
  At least one of the following **must** be true:
  
  - The [`extendedDynamicState3RepresentativeFragmentTestEnable`](#features-extendedDynamicState3RepresentativeFragmentTestEnable) feature is enabled
  - The [`shaderObject`](#features-shaderObject) feature is enabled

Valid Usage (Implicit)

- [](#VUID-vkCmdSetRepresentativeFragmentTestEnableNV-commandBuffer-parameter)VUID-vkCmdSetRepresentativeFragmentTestEnableNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdSetRepresentativeFragmentTestEnableNV-commandBuffer-recording)VUID-vkCmdSetRepresentativeFragmentTestEnableNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdSetRepresentativeFragmentTestEnableNV-commandBuffer-cmdpool)VUID-vkCmdSetRepresentativeFragmentTestEnableNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics operations
- [](#VUID-vkCmdSetRepresentativeFragmentTestEnableNV-videocoding)VUID-vkCmdSetRepresentativeFragmentTestEnableNV-videocoding  
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

vkCmdSetRepresentativeFragmentTestEnableNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_extended\_dynamic\_state3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_extended_dynamic_state3.html), [VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VK\_NV\_representative\_fragment\_test](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_representative_fragment_test.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdSetRepresentativeFragmentTestEnableNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
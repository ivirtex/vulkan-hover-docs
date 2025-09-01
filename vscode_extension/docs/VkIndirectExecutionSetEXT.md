# VkIndirectExecutionSetEXT(3) Manual Page

## Name

VkIndirectExecutionSetEXT - Opaque handle to an indirect execution set object



## [](#_c_specification)C Specification

Indirect Execution Sets contain sets of pipelines or shader objects which can be bound individually.

```c++
// Provided by VK_EXT_device_generated_commands
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkIndirectExecutionSetEXT)
```

## [](#_description)Description

Indirect Execution Sets allow the device to bind different shaders and pipeline states using [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#device-generated-commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#device-generated-commands).

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkGeneratedCommandsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoEXT.html), [VkGeneratedCommandsMemoryRequirementsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsMemoryRequirementsInfoEXT.html), [vkCreateIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateIndirectExecutionSetEXT.html), [vkDestroyIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyIndirectExecutionSetEXT.html), [vkUpdateIndirectExecutionSetPipelineEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateIndirectExecutionSetPipelineEXT.html), [vkUpdateIndirectExecutionSetShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateIndirectExecutionSetShaderEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectExecutionSetEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
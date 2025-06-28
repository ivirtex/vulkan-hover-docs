# VkIndirectExecutionSetInfoTypeEXT(3) Manual Page

## Name

VkIndirectExecutionSetInfoTypeEXT - Enum specifying allowed usage of an indirect execution set



## [](#_c_specification)C Specification

Values which **can** be set in [VkIndirectExecutionSetCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetCreateInfoEXT.html)::`type`, specifying contents of an indirect execution set, are:

```c++
// Provided by VK_EXT_device_generated_commands
typedef enum VkIndirectExecutionSetInfoTypeEXT {
    VK_INDIRECT_EXECUTION_SET_INFO_TYPE_PIPELINES_EXT = 0,
    VK_INDIRECT_EXECUTION_SET_INFO_TYPE_SHADER_OBJECTS_EXT = 1,
} VkIndirectExecutionSetInfoTypeEXT;
```

## [](#_description)Description

- `VK_INDIRECT_EXECUTION_SET_INFO_TYPE_PIPELINES_EXT` specifies that the indirect execution set contains [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) objects.
- `VK_INDIRECT_EXECUTION_SET_INFO_TYPE_SHADER_OBJECTS_EXT` specifies that the indirect execution set contains [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) objects.

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectCommandsExecutionSetTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsExecutionSetTokenEXT.html), [VkIndirectExecutionSetCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetCreateInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectExecutionSetInfoTypeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
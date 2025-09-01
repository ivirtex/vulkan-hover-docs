# VkIndirectExecutionSetInfoEXT(3) Manual Page

## Name

VkIndirectExecutionSetInfoEXT - Union specifying parameters of a newly created indirect execution set



## [](#_c_specification)C Specification

The `VkIndirectExecutionSetInfoEXT` union is defined as:

```c++
// Provided by VK_EXT_device_generated_commands
typedef union VkIndirectExecutionSetInfoEXT {
    const VkIndirectExecutionSetPipelineInfoEXT*    pPipelineInfo;
    const VkIndirectExecutionSetShaderInfoEXT*      pShaderInfo;
} VkIndirectExecutionSetInfoEXT;
```

## [](#_members)Members

- `pPipelineInfo` is a pointer to a [VkIndirectExecutionSetPipelineInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetPipelineInfoEXT.html) structure containing pipeline layout information for the set.
- `pShaderInfo` is a pointer to a [VkIndirectExecutionSetShaderInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetShaderInfoEXT.html) structure containing shader object layout information for the set.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectExecutionSetCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetCreateInfoEXT.html), [VkIndirectExecutionSetPipelineInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetPipelineInfoEXT.html), [VkIndirectExecutionSetShaderInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetShaderInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectExecutionSetInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VkIndirectExecutionSetCreateInfoEXT(3) Manual Page

## Name

VkIndirectExecutionSetCreateInfoEXT - Structure specifying parameters of a newly created indirect execution set



## [](#_c_specification)C Specification

The `VkIndirectExecutionSetCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkIndirectExecutionSetCreateInfoEXT {
    VkStructureType                      sType;
    const void*                          pNext;
    VkIndirectExecutionSetInfoTypeEXT    type;
    VkIndirectExecutionSetInfoEXT        info;
} VkIndirectExecutionSetCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `type` is a [VkIndirectExecutionSetInfoTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetInfoTypeEXT.html) describing the type of set being created and determining which field of the `info` union will be used.
- `info` is a [VkIndirectExecutionSetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetInfoEXT.html) union containing layout information for the set.

## [](#_description)Description

Valid Usage

- [](#VUID-VkIndirectExecutionSetCreateInfoEXT-maxIndirectShaderObjectCount-11014)VUID-VkIndirectExecutionSetCreateInfoEXT-maxIndirectShaderObjectCount-11014  
  If `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT`::`maxIndirectShaderObjectCount` is zero or the [`shaderObject`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderObject) feature is not enabled `type` **must** not be `VK_INDIRECT_EXECUTION_SET_INFO_TYPE_SHADER_OBJECTS_EXT`

Valid Usage (Implicit)

- [](#VUID-VkIndirectExecutionSetCreateInfoEXT-sType-sType)VUID-VkIndirectExecutionSetCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_INDIRECT_EXECUTION_SET_CREATE_INFO_EXT`
- [](#VUID-VkIndirectExecutionSetCreateInfoEXT-type-parameter)VUID-VkIndirectExecutionSetCreateInfoEXT-type-parameter  
  `type` **must** be a valid [VkIndirectExecutionSetInfoTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetInfoTypeEXT.html) value
- [](#VUID-VkIndirectExecutionSetCreateInfoEXT-pPipelineInfo-parameter)VUID-VkIndirectExecutionSetCreateInfoEXT-pPipelineInfo-parameter  
  If `type` is `VK_INDIRECT_EXECUTION_SET_INFO_TYPE_PIPELINES_EXT`, the `pPipelineInfo` member of `info` **must** be a valid pointer to a valid [VkIndirectExecutionSetPipelineInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetPipelineInfoEXT.html) structure
- [](#VUID-VkIndirectExecutionSetCreateInfoEXT-pShaderInfo-parameter)VUID-VkIndirectExecutionSetCreateInfoEXT-pShaderInfo-parameter  
  If `type` is `VK_INDIRECT_EXECUTION_SET_INFO_TYPE_SHADER_OBJECTS_EXT`, the `pShaderInfo` member of `info` **must** be a valid pointer to a valid [VkIndirectExecutionSetShaderInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetShaderInfoEXT.html) structure

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectExecutionSetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetInfoEXT.html), [VkIndirectExecutionSetInfoTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetInfoTypeEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateIndirectExecutionSetEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectExecutionSetCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
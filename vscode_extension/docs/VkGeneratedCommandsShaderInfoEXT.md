# VkGeneratedCommandsShaderInfoEXT(3) Manual Page

## Name

VkGeneratedCommandsShaderInfoEXT - Structure specifying shader objects for use with indirect command preprocessing



## [](#_c_specification)C Specification

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkGeneratedCommandsShaderInfoEXT {
    VkStructureType       sType;
    void*                 pNext;
    uint32_t              shaderCount;
    const VkShaderEXT*    pShaders;
} VkGeneratedCommandsShaderInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `shaderCount` is the size of the `pShaders` array.
- `pShaders` is a pointer to an array of shader objects.

## [](#_description)Description

Valid Usage

- [](#VUID-VkGeneratedCommandsShaderInfoEXT-pShaders-11127)VUID-VkGeneratedCommandsShaderInfoEXT-pShaders-11127  
  `pShaders` **must** not contain more than one shader object for a given [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) stage

Valid Usage (Implicit)

- [](#VUID-VkGeneratedCommandsShaderInfoEXT-sType-sType)VUID-VkGeneratedCommandsShaderInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GENERATED_COMMANDS_SHADER_INFO_EXT`
- [](#VUID-VkGeneratedCommandsShaderInfoEXT-pShaders-parameter)VUID-VkGeneratedCommandsShaderInfoEXT-pShaders-parameter  
  `pShaders` **must** be a valid pointer to an array of `shaderCount` valid [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) handles
- [](#VUID-VkGeneratedCommandsShaderInfoEXT-shaderCount-arraylength)VUID-VkGeneratedCommandsShaderInfoEXT-shaderCount-arraylength  
  `shaderCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkGeneratedCommandsShaderInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
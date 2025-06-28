# VkPipelineShaderStageModuleIdentifierCreateInfoEXT(3) Manual Page

## Name

VkPipelineShaderStageModuleIdentifierCreateInfoEXT - Structure specifying an identifier for a shader module



## [](#_c_specification)C Specification

An identifier **can** be provided instead of shader code in an attempt to compile pipelines without providing complete SPIR-V to the implementation.

The `VkPipelineShaderStageModuleIdentifierCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_shader_module_identifier
typedef struct VkPipelineShaderStageModuleIdentifierCreateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           identifierSize;
    const uint8_t*     pIdentifier;
} VkPipelineShaderStageModuleIdentifierCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `identifierSize` is the size, in bytes, of the buffer pointed to by `pIdentifier`.
- `pIdentifier` is a pointer to a buffer of opaque data specifying an identifier.

## [](#_description)Description

Any identifier **can** be used. If the pipeline being created with identifier requires compilation to complete the pipeline creation call, pipeline compilation **must** fail as defined by `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT`.

`pIdentifier` and `identifierSize` **can** be obtained from an [VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleIdentifierEXT.html) queried earlier.

Valid Usage

- [](#VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-pNext-06850)VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-pNext-06850  
  If this structure is included in a `pNext` chain and `identifierSize` is not equal to 0, the [`shaderModuleIdentifier`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderModuleIdentifier) feature **must** be enabled
- [](#VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-pNext-06851)VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-pNext-06851  
  If this structure is included in a `pNext` chain of [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html) and `identifierSize` is not equal to 0, the pipeline **must** be created with the `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` flag set
- [](#VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-identifierSize-06852)VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-identifierSize-06852  
  `identifierSize` **must** be less-or-equal to `VK_MAX_SHADER_MODULE_IDENTIFIER_SIZE_EXT`

Valid Usage (Implicit)

- [](#VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-sType-sType)VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_MODULE_IDENTIFIER_CREATE_INFO_EXT`
- [](#VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-pIdentifier-parameter)VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-pIdentifier-parameter  
  If `identifierSize` is not `0`, `pIdentifier` **must** be a valid pointer to an array of `identifierSize` `uint8_t` values

## [](#_see_also)See Also

[VK\_EXT\_shader\_module\_identifier](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_module_identifier.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineShaderStageModuleIdentifierCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
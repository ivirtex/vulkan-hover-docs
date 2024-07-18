# VkPipelineShaderStageModuleIdentifierCreateInfoEXT(3) Manual Page

## Name

VkPipelineShaderStageModuleIdentifierCreateInfoEXT - Structure
specifying an identifier for a shader module



## <a href="#_c_specification" class="anchor"></a>C Specification

An identifier **can** be provided instead of shader code in an attempt
to compile pipelines without providing complete SPIR-V to the
implementation.

The `VkPipelineShaderStageModuleIdentifierCreateInfoEXT` structure is
defined as:

``` c
// Provided by VK_EXT_shader_module_identifier
typedef struct VkPipelineShaderStageModuleIdentifierCreateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           identifierSize;
    const uint8_t*     pIdentifier;
} VkPipelineShaderStageModuleIdentifierCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `identifierSize` is the size, in bytes, of the buffer pointed to by
  `pIdentifier`.

- `pIdentifier` is a pointer to a buffer of opaque data specifying an
  identifier.

## <a href="#_description" class="anchor"></a>Description

Any identifier **can** be used. If the pipeline being created with
identifier requires compilation to complete the pipeline creation call,
pipeline compilation **must** fail as defined by
`VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT`.

`pIdentifier` and `identifierSize` **can** be obtained from an
[VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleIdentifierEXT.html) queried
earlier.

Valid Usage

- <a
  href="#VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-pNext-06850"
  id="VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-pNext-06850"></a>
  VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-pNext-06850  
  If this structure is included in a `pNext` chain and `identifierSize`
  is not equal to 0, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderModuleIdentifier"
  target="_blank" rel="noopener"><code>shaderModuleIdentifier</code></a>
  feature **must** be enabled

- <a
  href="#VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-pNext-06851"
  id="VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-pNext-06851"></a>
  VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-pNext-06851  
  If this struct is included in a `pNext` chain of
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  and `identifierSize` is not equal to 0, the pipeline **must** be
  created with the
  `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` flag set

- <a
  href="#VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-identifierSize-06852"
  id="VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-identifierSize-06852"></a>
  VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-identifierSize-06852  
  `identifierSize` **must** be less-or-equal to
  `VK_MAX_SHADER_MODULE_IDENTIFIER_SIZE_EXT`

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-sType-sType"
  id="VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-sType-sType"></a>
  VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_MODULE_IDENTIFIER_CREATE_INFO_EXT`

- <a
  href="#VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-pIdentifier-parameter"
  id="VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-pIdentifier-parameter"></a>
  VUID-VkPipelineShaderStageModuleIdentifierCreateInfoEXT-pIdentifier-parameter  
  If `identifierSize` is not `0`, `pIdentifier` **must** be a valid
  pointer to an array of `identifierSize` `uint8_t` values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_module_identifier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_module_identifier.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineShaderStageModuleIdentifierCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

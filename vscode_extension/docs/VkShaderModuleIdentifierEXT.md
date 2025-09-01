# VkShaderModuleIdentifierEXT(3) Manual Page

## Name

VkShaderModuleIdentifierEXT - A unique identifier for a shader module



## [](#_c_specification)C Specification

[VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleIdentifierEXT.html) represents a shader module identifier returned by the implementation.

```c++
// Provided by VK_EXT_shader_module_identifier
typedef struct VkShaderModuleIdentifierEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           identifierSize;
    uint8_t            identifier[VK_MAX_SHADER_MODULE_IDENTIFIER_SIZE_EXT];
} VkShaderModuleIdentifierEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `identifierSize` is the size, in bytes, of valid data returned in `identifier`.
- `identifier` is a buffer of opaque data specifying an identifier.

## [](#_description)Description

Any returned values beyond the first `identifierSize` bytes are undefined. Implementations **must** return an `identifierSize` greater than 0, and less-or-equal to `VK_MAX_SHADER_MODULE_IDENTIFIER_SIZE_EXT`.

Two identifiers are considered equal if `identifierSize` is equal and the first `identifierSize` bytes of `identifier` compare equal.

Implementations **may** return a different `identifierSize` for different modules. Implementations **should** ensure that `identifierSize` is large enough to uniquely define a shader module.

Valid Usage (Implicit)

- [](#VUID-VkShaderModuleIdentifierEXT-sType-sType)VUID-VkShaderModuleIdentifierEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SHADER_MODULE_IDENTIFIER_EXT`
- [](#VUID-VkShaderModuleIdentifierEXT-pNext-pNext)VUID-VkShaderModuleIdentifierEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_EXT\_shader\_module\_identifier](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_module_identifier.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetShaderModuleCreateInfoIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetShaderModuleCreateInfoIdentifierEXT.html), [vkGetShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetShaderModuleIdentifierEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkShaderModuleIdentifierEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
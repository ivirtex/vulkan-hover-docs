# VkShaderModuleIdentifierEXT(3) Manual Page

## Name

VkShaderModuleIdentifierEXT - A unique identifier for a shader module



## <a href="#_c_specification" class="anchor"></a>C Specification

[VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleIdentifierEXT.html)
represents a shader module identifier returned by the implementation.

``` c
// Provided by VK_EXT_shader_module_identifier
typedef struct VkShaderModuleIdentifierEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           identifierSize;
    uint8_t            identifier[VK_MAX_SHADER_MODULE_IDENTIFIER_SIZE_EXT];
} VkShaderModuleIdentifierEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `identifierSize` is the size, in bytes, of valid data returned in
  `identifier`.

- `identifier` is a buffer of opaque data specifying an identifier.

## <a href="#_description" class="anchor"></a>Description

Any returned values beyond the first `identifierSize` bytes are
undefined. Implementations **must** return an `identifierSize` greater
than 0, and less-or-equal to `VK_MAX_SHADER_MODULE_IDENTIFIER_SIZE_EXT`.

Two identifiers are considered equal if `identifierSize` is equal and
the first `identifierSize` bytes of `identifier` compare equal.

Implementations **may** return a different `identifierSize` for
different modules. Implementations **should** ensure that
`identifierSize` is large enough to uniquely define a shader module.

Valid Usage (Implicit)

- <a href="#VUID-VkShaderModuleIdentifierEXT-sType-sType"
  id="VUID-VkShaderModuleIdentifierEXT-sType-sType"></a>
  VUID-VkShaderModuleIdentifierEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SHADER_MODULE_IDENTIFIER_EXT`

- <a href="#VUID-VkShaderModuleIdentifierEXT-pNext-pNext"
  id="VUID-VkShaderModuleIdentifierEXT-pNext-pNext"></a>
  VUID-VkShaderModuleIdentifierEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_module_identifier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_module_identifier.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetShaderModuleCreateInfoIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetShaderModuleCreateInfoIdentifierEXT.html),
[vkGetShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetShaderModuleIdentifierEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkShaderModuleIdentifierEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

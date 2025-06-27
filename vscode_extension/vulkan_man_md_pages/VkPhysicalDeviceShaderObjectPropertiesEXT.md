# VkPhysicalDeviceShaderObjectPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceShaderObjectPropertiesEXT - Structure describing shader
object properties supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderObjectPropertiesEXT` structure is defined as:

``` c
// Provided by VK_EXT_shader_object
typedef struct VkPhysicalDeviceShaderObjectPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint8_t            shaderBinaryUUID[VK_UUID_SIZE];
    uint32_t           shaderBinaryVersion;
} VkPhysicalDeviceShaderObjectPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-shaderBinaryUUID"></span> `shaderBinaryUUID` is an
  array of `VK_UUID_SIZE` `uint8_t` values representing a universally
  unique identifier for one or more implementations whose shader
  binaries are guaranteed to be compatible with each other.

- <span id="limits-shaderBinaryVersion"></span> `shaderBinaryVersion` is
  an unsigned integer incremented to represent backwards compatible
  differences between implementations with the same `shaderBinaryUUID`.

## <a href="#_description" class="anchor"></a>Description

The purpose and usage of the values of this structure are described in
greater detail in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects-binary-compatibility"
target="_blank" rel="noopener">Binary Shader Compatibility</a>.

If the `VkPhysicalDeviceShaderObjectPropertiesEXT` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceShaderObjectPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceShaderObjectPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderObjectPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_OBJECT_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderObjectPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

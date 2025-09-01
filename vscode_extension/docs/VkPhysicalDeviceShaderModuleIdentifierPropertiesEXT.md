# VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT - Structure describing shader module identifier properties of an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_shader_module_identifier
typedef struct VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint8_t            shaderModuleIdentifierAlgorithmUUID[VK_UUID_SIZE];
} VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT;
```

## [](#_members)Members

The members of the `VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT` structure describe the following:

## [](#_description)Description

- []()`shaderModuleIdentifierAlgorithmUUID` is an array of `VK_UUID_SIZE` `uint8_t` values which uniquely represents the algorithm used to compute an identifier in [vkGetShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetShaderModuleIdentifierEXT.html) and [vkGetShaderModuleCreateInfoIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetShaderModuleCreateInfoIdentifierEXT.html). Implementations **should** not change this value in different driver versions if the algorithm used to compute an identifier is the same.

Note

The algorithm UUID may be the same in different ICDs if the algorithms are guaranteed to produce the same results. This may happen in driver stacks which support different kinds of hardware with shared code.

Khronos' conformance testing can not guarantee that `shaderModuleIdentifierAlgorithmUUID` values are actually unique, so implementors should make their own best efforts to ensure that their UUID is unlikely to conflict with other implementations which may use a different algorithm. In particular, hard-coded values which easily conflict, such as all-`0` bits, **should** never be used. Hard-coded values are acceptable if best effort is ensured that the value will not accidentally conflict.

If the `VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_MODULE_IDENTIFIER_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_shader\_module\_identifier](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_module_identifier.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# vkGetShaderModuleCreateInfoIdentifierEXT(3) Manual Page

## Name

vkGetShaderModuleCreateInfoIdentifierEXT - Query a unique identifier for a shader module create info



## [](#_c_specification)C Specification

[VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) structures have unique identifiers associated with them. To query an implementation provided identifier, call:

```c++
// Provided by VK_EXT_shader_module_identifier
void vkGetShaderModuleCreateInfoIdentifierEXT(
    VkDevice                                    device,
    const VkShaderModuleCreateInfo*             pCreateInfo,
    VkShaderModuleIdentifierEXT*                pIdentifier);
```

## [](#_parameters)Parameters

- `device` is the logical device that **can** create a [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html) from `pCreateInfo`.
- `pCreateInfo` is a pointer to a [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) structure.
- `pIdentifier` is a pointer to the returned [VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleIdentifierEXT.html).

## [](#_description)Description

The identifier returned by implementation **must** only depend on `shaderIdentifierAlgorithmUUID` and information provided in the [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html). The implementation **may** return equal identifiers for two different [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) structures if the difference does not affect pipeline compilation. Identifiers are only meaningful on different [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) objects if the device the identifier was queried from had the same [`shaderModuleIdentifierAlgorithmUUID`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-shaderModuleIdentifierAlgorithmUUID) as the device consuming the identifier.

The identifier returned by the implementation in [vkGetShaderModuleCreateInfoIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetShaderModuleCreateInfoIdentifierEXT.html) **must** be equal to the identifier returned by [vkGetShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetShaderModuleIdentifierEXT.html) given equivalent definitions of [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) and any chained `pNext` structures.

Valid Usage

- [](#VUID-vkGetShaderModuleCreateInfoIdentifierEXT-shaderModuleIdentifier-06885)VUID-vkGetShaderModuleCreateInfoIdentifierEXT-shaderModuleIdentifier-06885  
  [`shaderModuleIdentifier`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderModuleIdentifier) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetShaderModuleCreateInfoIdentifierEXT-device-parameter)VUID-vkGetShaderModuleCreateInfoIdentifierEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetShaderModuleCreateInfoIdentifierEXT-pCreateInfo-parameter)VUID-vkGetShaderModuleCreateInfoIdentifierEXT-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) structure
- [](#VUID-vkGetShaderModuleCreateInfoIdentifierEXT-pIdentifier-parameter)VUID-vkGetShaderModuleCreateInfoIdentifierEXT-pIdentifier-parameter  
  `pIdentifier` **must** be a valid pointer to a [VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleIdentifierEXT.html) structure

## [](#_see_also)See Also

[VK\_EXT\_shader\_module\_identifier](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_module_identifier.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html), [VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleIdentifierEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetShaderModuleCreateInfoIdentifierEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
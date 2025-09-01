# vkGetShaderModuleIdentifierEXT(3) Manual Page

## Name

vkGetShaderModuleIdentifierEXT - Query a unique identifier for a shader module



## [](#_c_specification)C Specification

Shader modules have unique identifiers associated with them. To query an implementation provided identifier, call:

```c++
// Provided by VK_EXT_shader_module_identifier
void vkGetShaderModuleIdentifierEXT(
    VkDevice                                    device,
    VkShaderModule                              shaderModule,
    VkShaderModuleIdentifierEXT*                pIdentifier);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the shader module.
- `shaderModule` is the handle of the shader module.
- `pIdentifier` is a pointer to the returned [VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleIdentifierEXT.html).

## [](#_description)Description

The identifier returned by the implementation **must** only depend on `shaderIdentifierAlgorithmUUID` and information provided in the [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) which created `shaderModule`. The implementation **may** return equal identifiers for two different [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) structures if the difference does not affect pipeline compilation. Identifiers are only meaningful on different [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) objects if the device the identifier was queried from had the same [`shaderModuleIdentifierAlgorithmUUID`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-shaderModuleIdentifierAlgorithmUUID) as the device consuming the identifier.

Valid Usage

- [](#VUID-vkGetShaderModuleIdentifierEXT-shaderModuleIdentifier-06884)VUID-vkGetShaderModuleIdentifierEXT-shaderModuleIdentifier-06884  
  [`shaderModuleIdentifier`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderModuleIdentifier) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetShaderModuleIdentifierEXT-device-parameter)VUID-vkGetShaderModuleIdentifierEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetShaderModuleIdentifierEXT-shaderModule-parameter)VUID-vkGetShaderModuleIdentifierEXT-shaderModule-parameter  
  `shaderModule` **must** be a valid [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html) handle
- [](#VUID-vkGetShaderModuleIdentifierEXT-pIdentifier-parameter)VUID-vkGetShaderModuleIdentifierEXT-pIdentifier-parameter  
  `pIdentifier` **must** be a valid pointer to a [VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleIdentifierEXT.html) structure
- [](#VUID-vkGetShaderModuleIdentifierEXT-shaderModule-parent)VUID-vkGetShaderModuleIdentifierEXT-shaderModule-parent  
  `shaderModule` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_EXT\_shader\_module\_identifier](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_module_identifier.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html), [VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleIdentifierEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetShaderModuleIdentifierEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
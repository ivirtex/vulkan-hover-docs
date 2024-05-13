# vkGetShaderModuleIdentifierEXT(3) Manual Page

## Name

vkGetShaderModuleIdentifierEXT - Query a unique identifier for a shader
module



## <a href="#_c_specification" class="anchor"></a>C Specification

Shader modules have unique identifiers associated with them. To query an
implementation provided identifier, call:

``` c
// Provided by VK_EXT_shader_module_identifier
void vkGetShaderModuleIdentifierEXT(
    VkDevice                                    device,
    VkShaderModule                              shaderModule,
    VkShaderModuleIdentifierEXT*                pIdentifier);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that created the shader module.

- `shaderModule` is the handle of the shader module.

- `pIdentifier` is a pointer to the returned
  [VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleIdentifierEXT.html).

## <a href="#_description" class="anchor"></a>Description

The identifier returned by the implementation **must** only depend on
`shaderIdentifierAlgorithmUUID` and information provided in the
[VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) which created
`shaderModule`. The implementation **may** return equal identifiers for
two different [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html)
structures if the difference does not affect pipeline compilation.
Identifiers are only meaningful on different [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)
objects if the device the identifier was queried from had the same <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderModuleIdentifierAlgorithmUUID"
target="_blank"
rel="noopener"><code>shaderModuleIdentifierAlgorithmUUID</code></a> as
the device consuming the identifier.

Valid Usage

- <a
  href="#VUID-vkGetShaderModuleIdentifierEXT-shaderModuleIdentifier-06884"
  id="VUID-vkGetShaderModuleIdentifierEXT-shaderModuleIdentifier-06884"></a>
  VUID-vkGetShaderModuleIdentifierEXT-shaderModuleIdentifier-06884  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderModuleIdentifier"
  target="_blank" rel="noopener"><code>shaderModuleIdentifier</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkGetShaderModuleIdentifierEXT-device-parameter"
  id="VUID-vkGetShaderModuleIdentifierEXT-device-parameter"></a>
  VUID-vkGetShaderModuleIdentifierEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetShaderModuleIdentifierEXT-shaderModule-parameter"
  id="VUID-vkGetShaderModuleIdentifierEXT-shaderModule-parameter"></a>
  VUID-vkGetShaderModuleIdentifierEXT-shaderModule-parameter  
  `shaderModule` **must** be a valid
  [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html) handle

- <a href="#VUID-vkGetShaderModuleIdentifierEXT-pIdentifier-parameter"
  id="VUID-vkGetShaderModuleIdentifierEXT-pIdentifier-parameter"></a>
  VUID-vkGetShaderModuleIdentifierEXT-pIdentifier-parameter  
  `pIdentifier` **must** be a valid pointer to a
  [VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleIdentifierEXT.html)
  structure

- <a href="#VUID-vkGetShaderModuleIdentifierEXT-shaderModule-parent"
  id="VUID-vkGetShaderModuleIdentifierEXT-shaderModule-parent"></a>
  VUID-vkGetShaderModuleIdentifierEXT-shaderModule-parent  
  `shaderModule` **must** have been created, allocated, or retrieved
  from `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_module_identifier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_module_identifier.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html),
[VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleIdentifierEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetShaderModuleIdentifierEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

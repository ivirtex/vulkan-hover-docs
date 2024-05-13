# vkGetShaderModuleCreateInfoIdentifierEXT(3) Manual Page

## Name

vkGetShaderModuleCreateInfoIdentifierEXT - Query a unique identifier for
a shader module create info



## <a href="#_c_specification" class="anchor"></a>C Specification

[VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) structures
have unique identifiers associated with them. To query an implementation
provided identifier, call:

``` c
// Provided by VK_EXT_shader_module_identifier
void vkGetShaderModuleCreateInfoIdentifierEXT(
    VkDevice                                    device,
    const VkShaderModuleCreateInfo*             pCreateInfo,
    VkShaderModuleIdentifierEXT*                pIdentifier);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that **can** create a
  [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html) from `pCreateInfo`.

- `pCreateInfo` is a pointer to a
  [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) structure.

- `pIdentifier` is a pointer to the returned
  [VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleIdentifierEXT.html).

## <a href="#_description" class="anchor"></a>Description

The identifier returned by implementation **must** only depend on
`shaderIdentifierAlgorithmUUID` and information provided in the
[VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html). The
implementation **may** return equal identifiers for two different
[VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) structures if
the difference does not affect pipeline compilation. Identifiers are
only meaningful on different [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) objects if the
device the identifier was queried from had the same <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderModuleIdentifierAlgorithmUUID"
target="_blank"
rel="noopener"><code>shaderModuleIdentifierAlgorithmUUID</code></a> as
the device consuming the identifier.

The identifier returned by the implementation in
[vkGetShaderModuleCreateInfoIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetShaderModuleCreateInfoIdentifierEXT.html)
**must** be equal to the identifier returned by
[vkGetShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetShaderModuleIdentifierEXT.html)
given equivalent definitions of
[VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) and any
chained `pNext` structures.

Valid Usage

- <a
  href="#VUID-vkGetShaderModuleCreateInfoIdentifierEXT-shaderModuleIdentifier-06885"
  id="VUID-vkGetShaderModuleCreateInfoIdentifierEXT-shaderModuleIdentifier-06885"></a>
  VUID-vkGetShaderModuleCreateInfoIdentifierEXT-shaderModuleIdentifier-06885  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderModuleIdentifier"
  target="_blank" rel="noopener"><code>shaderModuleIdentifier</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetShaderModuleCreateInfoIdentifierEXT-device-parameter"
  id="VUID-vkGetShaderModuleCreateInfoIdentifierEXT-device-parameter"></a>
  VUID-vkGetShaderModuleCreateInfoIdentifierEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetShaderModuleCreateInfoIdentifierEXT-pCreateInfo-parameter"
  id="VUID-vkGetShaderModuleCreateInfoIdentifierEXT-pCreateInfo-parameter"></a>
  VUID-vkGetShaderModuleCreateInfoIdentifierEXT-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) structure

- <a
  href="#VUID-vkGetShaderModuleCreateInfoIdentifierEXT-pIdentifier-parameter"
  id="VUID-vkGetShaderModuleCreateInfoIdentifierEXT-pIdentifier-parameter"></a>
  VUID-vkGetShaderModuleCreateInfoIdentifierEXT-pIdentifier-parameter  
  `pIdentifier` **must** be a valid pointer to a
  [VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleIdentifierEXT.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_module_identifier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_module_identifier.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html),
[VkShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleIdentifierEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetShaderModuleCreateInfoIdentifierEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

# VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT - Structure
describing shader module identifier properties of an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_shader_module_identifier
typedef struct VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint8_t            shaderModuleIdentifierAlgorithmUUID[VK_UUID_SIZE];
} VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

The members of the `VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT`
structure describe the following:

## <a href="#_description" class="anchor"></a>Description

- <span id="limits-shaderModuleIdentifierAlgorithmUUID"></span>
  `shaderModuleIdentifierAlgorithmUUID` is an array of `VK_UUID_SIZE`
  `uint8_t` values which uniquely represents the algorithm used to
  compute an identifier in
  [vkGetShaderModuleIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetShaderModuleIdentifierEXT.html)
  and
  [vkGetShaderModuleCreateInfoIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetShaderModuleCreateInfoIdentifierEXT.html).
  Implementations **should** not change this value in different driver
  versions if the algorithm used to compute an identifier is the same.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The algorithm UUID may be the same in different ICDs if the
algorithms are guaranteed to produce the same results. This may happen
in driver stacks which support different kinds of hardware with shared
code.</p>
<p>Khronos' conformance testing can not guarantee that
<code>shaderModuleIdentifierAlgorithmUUID</code> values are actually
unique, so implementors should make their own best efforts to ensure
that their UUID is unlikely to conflict with other implementations which
may use a different algorithm. In particular, hard-coded values which
easily conflict, such as all-<code>0</code> bits,
<strong>should</strong> never be used. Hard-coded values are acceptable
if best effort is ensured that the value will not accidentally
conflict.</p></td>
</tr>
</tbody>
</table>

If the `VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_MODULE_IDENTIFIER_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_module_identifier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_module_identifier.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

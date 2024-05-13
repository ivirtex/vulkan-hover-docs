# VK_AMD_shader_core_properties2(3) Manual Page

## Name

VK_AMD_shader_core_properties2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

228

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_AMD_shader_core_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_shader_core_properties.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Matthaeus G. Chajdas <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_shader_core_properties2%5D%20@anteru%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_shader_core_properties2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>anteru</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-07-26

**IP Status**  
No known IP claims.

**Contributors**  
- Matthaeus G. Chajdas, AMD

- Tobias Hector, AMD

## <a href="#_description" class="anchor"></a>Description

This extension exposes additional shader core properties for a target
physical device through the
[`VK_KHR_get_physical_device_properties2`](VK_KHR_get_physical_device_properties2.html)
extension.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceShaderCoreProperties2AMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderCoreProperties2AMD.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkShaderCorePropertiesFlagBitsAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCorePropertiesFlagBitsAMD.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkShaderCorePropertiesFlagsAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCorePropertiesFlagsAMD.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMD_SHADER_CORE_PROPERTIES_2_EXTENSION_NAME`

- `VK_AMD_SHADER_CORE_PROPERTIES_2_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_PROPERTIES_2_AMD`

## <a href="#_examples" class="anchor"></a>Examples

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-07-26 (Matthaeus G. Chajdas)

  - Initial draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMD_shader_core_properties2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

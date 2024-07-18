# VK_EXT_shader_replicated_composites(3) Manual Page

## Name

VK_EXT_shader_replicated_composites - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

565

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_EXT_replicated_composites](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/EXT/SPV_EXT_replicated_composites.html)

## <a href="#_contact" class="anchor"></a>Contact

- Kevin Petit <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_replicated_composites%5D%20@kpet%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_replicated_composites%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>kpet</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_shader_replicated_composites](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_shader_replicated_composites.adoc)

**Last Modified Date**  
2024-02-08

**IP Status**  
No known IP claims.

**Contributors**  
- Kévin Petit, Arm Ltd.

- Jeff Bolz, NVIDIA

- Piers Daniell, NVIDIA

This extension adds support for creating composites from a single value
in SPIR-V modules, as defined by SPV_EXT_replicated_composites.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_SHADER_REPLICATED_COMPOSITES_EXTENSION_NAME`

- `VK_EXT_SHADER_REPLICATED_COMPOSITES_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_REPLICATED_COMPOSITES_FEATURES_EXT`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-ReplicatedCompositesEXT"
  target="_blank" rel="noopener">ReplicatedCompositesEXT</a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2024-02-08 (Kévin Petit)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_shader_replicated_composites"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

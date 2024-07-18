# VK_VALVE_descriptor_set_host_mapping(3) Manual Page

## Name

VK_VALVE_descriptor_set_host_mapping - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

421

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">D3D support</a>

## <a href="#_contact" class="anchor"></a>Contact

- Hans-Kristian Arntzen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_VALVE_descriptor_set_host_mapping%5D%20@HansKristian-Work%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_VALVE_descriptor_set_host_mapping%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>HansKristian-Work</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-02-22

**IP Status**  
No known IP claims.

**Contributors**  
- Hans-Kristian Arntzen, Valve

## <a href="#_description" class="anchor"></a>Description

This extension allows applications to directly query a host pointer for
a [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html) which **can** be used to copy
descriptors between descriptor sets without the use of an API command.
Memory offsets and sizes for descriptors **can** be queried from a
[VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) as well.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>There is currently no specification language written for this
extension. The links to APIs defined by the extension are to stubs that
only include generated content such as API declarations and implicit
valid usage statements.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>This extension is only intended for use in specific embedded
environments with known implementation details, and is therefore
undocumented.</p></td>
</tr>
</tbody>
</table>

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetDescriptorSetHostMappingVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetHostMappingVALVE.html)

- [vkGetDescriptorSetLayoutHostMappingInfoVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutHostMappingInfoVALVE.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkDescriptorSetBindingReferenceVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetBindingReferenceVALVE.html)

- [VkDescriptorSetLayoutHostMappingInfoVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutHostMappingInfoVALVE.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceDescriptorSetHostMappingFeaturesVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorSetHostMappingFeaturesVALVE.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_VALVE_DESCRIPTOR_SET_HOST_MAPPING_EXTENSION_NAME`

- `VK_VALVE_DESCRIPTOR_SET_HOST_MAPPING_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_BINDING_REFERENCE_VALVE`

  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_HOST_MAPPING_INFO_VALVE`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_SET_HOST_MAPPING_FEATURES_VALVE`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-02-22 (Hans-Kristian Arntzen)

  - Initial specification

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_VALVE_descriptor_set_host_mapping"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

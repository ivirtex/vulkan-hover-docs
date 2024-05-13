# VK_VALVE_mutable_descriptor_type(3) Manual Page

## Name

VK_VALVE_mutable_descriptor_type - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

352

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_maintenance3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance3.html)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to
  [VK_EXT_mutable_descriptor_type](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_mutable_descriptor_type.html)
  extension

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">D3D support</a>

## <a href="#_contact" class="anchor"></a>Contact

- Joshua Ashton <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_VALVE_mutable_descriptor_type%5D%20@Joshua-Ashton%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_VALVE_mutable_descriptor_type%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>Joshua-Ashton</a>

- Hans-Kristian Arntzen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_VALVE_mutable_descriptor_type%5D%20@HansKristian-Work%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_VALVE_mutable_descriptor_type%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>HansKristian-Work</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-12-02

**IP Status**  
No known IP claims.

**Contributors**  
- Joshua Ashton, Valve

- Hans-Kristian Arntzen, Valve

## <a href="#_description" class="anchor"></a>Description

This extension allows applications to reduce descriptor memory footprint
by allowing a descriptor to be able to mutate to a given list of
descriptor types depending on which descriptor types are written into,
or copied into a descriptor set.

The main use case this extension intends to address is descriptor
indexing with `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`
where the descriptor types are completely generic, as this means
applications can allocate one large descriptor set, rather than having
one large descriptor set per descriptor type, which significantly bloats
descriptor memory usage and causes performance issues.

This extension also adds a mechanism to declare that a descriptor pool,
and therefore the descriptor sets that are allocated from it, reside
only in host memory; as such these descriptors can only be
updated/copied, but not bound.

These features together allow much more efficient emulation of the raw
D3D12 binding model. This extension is primarily intended to be useful
for API layering efforts.

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkMutableDescriptorTypeListVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeListVALVE.html)

- Extending
  [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html),
  [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateInfo.html):

  - [VkMutableDescriptorTypeCreateInfoVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeCreateInfoVALVE.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceMutableDescriptorTypeFeaturesVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMutableDescriptorTypeFeaturesVALVE.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_VALVE_MUTABLE_DESCRIPTOR_TYPE_EXTENSION_NAME`

- `VK_VALVE_MUTABLE_DESCRIPTOR_TYPE_SPEC_VERSION`

- Extending
  [VkDescriptorPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateFlagBits.html):

  - `VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_VALVE`

- Extending
  [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateFlagBits.html):

  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_VALVE`

- Extending [VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html):

  - `VK_DESCRIPTOR_TYPE_MUTABLE_VALVE`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_MUTABLE_DESCRIPTOR_TYPE_CREATE_INFO_VALVE`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MUTABLE_DESCRIPTOR_TYPE_FEATURES_VALVE`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-12-01 (Joshua Ashton, Hans-Kristian Arntzen)

  - Initial specification, squashed from public draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_VALVE_mutable_descriptor_type"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

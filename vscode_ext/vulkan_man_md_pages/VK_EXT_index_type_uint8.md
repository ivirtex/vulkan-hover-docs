# VK_EXT_index_type_uint8(3) Manual Page

## Name

VK_EXT_index_type_uint8 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

266

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to [VK_KHR_index_type_uint8](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_index_type_uint8.html)
  extension

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_index_type_uint8%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_index_type_uint8%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-05-02

**IP Status**  
No known IP claims.

**Contributors**  
- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension allows `uint8_t` indices to be used with
[vkCmdBindIndexBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer.html).

## <a href="#_promotion_to_vk_khr_index_type_uint8" class="anchor"></a>Promotion to `VK_KHR_index_type_uint8`

All functionality in this extension is included in
[`VK_KHR_index_type_uint8`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_index_type_uint8.html), with the
suffix changed to KHR. The original enum names are still available as
aliases of the KHR functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceIndexTypeUint8FeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceIndexTypeUint8FeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_INDEX_TYPE_UINT8_EXTENSION_NAME`

- `VK_EXT_INDEX_TYPE_UINT8_SPEC_VERSION`

- Extending [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html):

  - `VK_INDEX_TYPE_UINT8_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INDEX_TYPE_UINT8_FEATURES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-05-02 (Piers Daniell)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_index_type_uint8"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

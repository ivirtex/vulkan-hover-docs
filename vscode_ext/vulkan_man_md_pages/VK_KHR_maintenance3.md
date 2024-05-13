# VK_KHR_maintenance3(3) Manual Page

## Name

VK_KHR_maintenance3 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

169

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_maintenance3%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_maintenance3%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-09-05

**Contributors**  
- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

`VK_KHR_maintenance3` adds a collection of minor features that were
intentionally left out or overlooked from the original Vulkan 1.0
release.

The new features are as follows:

- A limit on the maximum number of descriptors that are supported in a
  single descriptor set layout. Some implementations have a limit on the
  total size of descriptors in a set, which cannot be expressed in terms
  of the limits in Vulkan 1.0.

- A limit on the maximum size of a single memory allocation. Some
  platforms have kernel interfaces that limit the maximum size of an
  allocation.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetDescriptorSetLayoutSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutSupportKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkDescriptorSetLayoutSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutSupportKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceMaintenance3PropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance3PropertiesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_MAINTENANCE3_EXTENSION_NAME`

- `VK_KHR_MAINTENANCE3_SPEC_VERSION`

- `VK_KHR_MAINTENANCE_3_EXTENSION_NAME`

- `VK_KHR_MAINTENANCE_3_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-08-22

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_maintenance3"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

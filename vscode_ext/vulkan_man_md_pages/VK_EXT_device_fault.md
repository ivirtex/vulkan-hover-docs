# VK_EXT_device_fault(3) Manual Page

## Name

VK_EXT_device_fault - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

342

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Ralph Potter r_potter

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_device_fault](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_device_fault.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-03-10

**IP Status**  
No known IP claims.

**Contributors**  
- Ralph Potter, Samsung

- Stuart Smith, AMD

- Jan-Harald Fredriksen, ARM

- Mark Bellamy, ARM

- Andrew Ellem, Google

- Alex Walters, IMG

- Jeff Bolz, NVIDIA

- Baldur Karlsson, Valve

## <a href="#_description" class="anchor"></a>Description

Device loss can be triggered by a variety of issues, including invalid
API usage, implementation errors, or hardware failures.

This extension introduces a new command:
[vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceFaultInfoEXT.html), which may be
called subsequent to a `VK_ERROR_DEVICE_LOST` error code having been
returned by the implementation. This command allows developers to query
for additional information on GPU faults which may have caused device
loss, and to generate binary crash dumps, which may be loaded into
external tools for further diagnosis.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceFaultInfoEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultAddressInfoEXT.html)

- [VkDeviceFaultCountsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultCountsEXT.html)

- [VkDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultInfoEXT.html)

- [VkDeviceFaultVendorBinaryHeaderVersionOneEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultVendorBinaryHeaderVersionOneEXT.html)

- [VkDeviceFaultVendorInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultVendorInfoEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceFaultFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFaultFeaturesEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkDeviceFaultAddressTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultAddressTypeEXT.html)

- [VkDeviceFaultVendorBinaryHeaderVersionEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultVendorBinaryHeaderVersionEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_DEVICE_FAULT_EXTENSION_NAME`

- `VK_EXT_DEVICE_FAULT_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DEVICE_FAULT_COUNTS_EXT`

  - `VK_STRUCTURE_TYPE_DEVICE_FAULT_INFO_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FAULT_FEATURES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2023-04-05 (Ralph Potter)

  - Restored two missing members to the XML definition of
    VkDeviceFaultVendorBinaryHeaderVersionOneEXT. No functional change
    to the specification.

- Revision 1, 2020-10-19 (Ralph Potter)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_device_fault"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

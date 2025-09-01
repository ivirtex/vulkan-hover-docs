# VK\_EXT\_pci\_bus\_info(3) Manual Page

## Name

VK\_EXT\_pci\_bus\_info - device extension



## [](#_registered_extension_number)Registered Extension Number

213

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Matthaeus G. Chajdas [\[GitHub\]anteru](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_pci_bus_info%5D%20%40anteru%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_pci_bus_info%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-12-10

**IP Status**

No known IP claims.

**Contributors**

- Matthaeus G. Chajdas, AMD
- Daniel Rakos, AMD

## [](#_description)Description

This extension adds a new query to obtain PCI bus information about a physical device.

Not all physical devices have PCI bus information, either due to the device not being connected to the system through a PCI interface or due to platform specific restrictions and policies. Thus this extension is only expected to be supported by physical devices which can provide the information.

As a consequence, applications should always check for the presence of the extension string for each individual physical device for which they intend to issue the new query for and should not have any assumptions about the availability of the extension on any given platform.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDevicePCIBusInfoPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePCIBusInfoPropertiesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_PCI_BUS_INFO_EXTENSION_NAME`
- `VK_EXT_PCI_BUS_INFO_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PCI_BUS_INFO_PROPERTIES_EXT`

## [](#_version_history)Version History

- Revision 2, 2018-12-10 (Daniel Rakos)
  
  - Changed all members of the new structure to have the uint32\_t type
- Revision 1, 2018-10-11 (Daniel Rakos)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_pci_bus_info).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
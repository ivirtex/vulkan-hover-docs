# VK_EXT_pci_bus_info(3) Manual Page

## Name

VK_EXT_pci_bus_info - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

213

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Matthaeus G. Chajdas <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_pci_bus_info%5D%20@anteru%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_pci_bus_info%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>anteru</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-12-10

**IP Status**  
No known IP claims.

**Contributors**  
- Matthaeus G. Chajdas, AMD

- Daniel Rakos, AMD

## <a href="#_description" class="anchor"></a>Description

This extension adds a new query to obtain PCI bus information about a
physical device.

Not all physical devices have PCI bus information, either due to the
device not being connected to the system through a PCI interface or due
to platform specific restrictions and policies. Thus this extension is
only expected to be supported by physical devices which can provide the
information.

As a consequence, applications should always check for the presence of
the extension string for each individual physical device for which they
intend to issue the new query for and should not have any assumptions
about the availability of the extension on any given platform.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDevicePCIBusInfoPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePCIBusInfoPropertiesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_PCI_BUS_INFO_EXTENSION_NAME`

- `VK_EXT_PCI_BUS_INFO_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PCI_BUS_INFO_PROPERTIES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2018-12-10 (Daniel Rakos)

  - Changed all members of the new structure to have the uint32_t type

- Revision 1, 2018-10-11 (Daniel Rakos)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_pci_bus_info"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

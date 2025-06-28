# VK\_KHR\_get\_surface\_capabilities2(3) Manual Page

## Name

VK\_KHR\_get\_surface\_capabilities2 - instance extension



## [](#_registered_extension_number)Registered Extension Number

120

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_get_surface_capabilities2%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_get_surface_capabilities2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-02-27

**IP Status**

No known IP claims.

**Contributors**

- Ian Elliott, Google
- James Jones, NVIDIA
- Alon Or-bach, Samsung

## [](#_description)Description

This extension provides new queries for device surface capabilities that can be easily extended by other extensions, without introducing any further queries. This extension can be considered the `VK_KHR_surface` equivalent of the `VK_KHR_get_physical_device_properties2` extension.

## [](#_new_commands)New Commands

- [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html)
- [vkGetPhysicalDeviceSurfaceFormats2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceFormats2KHR.html)

## [](#_new_structures)New Structures

- [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)
- [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html)
- [VkSurfaceFormat2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFormat2KHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_GET_SURFACE_CAPABILITIES_2_EXTENSION_NAME`
- `VK_KHR_GET_SURFACE_CAPABILITIES_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR`
  - `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_KHR`
  - `VK_STRUCTURE_TYPE_SURFACE_FORMAT_2_KHR`

## [](#_issues)Issues

1\) What should this extension be named?

**RESOLVED**: `VK_KHR_get_surface_capabilities2`. Other alternatives:

- `VK_KHR_surface2`
- One extension, combining a separate display-specific query extension.

2\) Should additional WSI query functions be extended?

**RESOLVED**:

- [vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html): Yes. The need for this motivated the extension.
- [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html): No. Currently only has boolean output. Extensions should instead extend [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html).
- [vkGetPhysicalDeviceSurfaceFormatsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html): Yes.
- [vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html): No. Recent discussion concluded this introduced too much variability for applications to deal with. Extensions should instead extend [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html).
- [vkGetPhysicalDeviceXlibPresentationSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceXlibPresentationSupportKHR.html): Not in this extension.
- [vkGetPhysicalDeviceXcbPresentationSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceXcbPresentationSupportKHR.html): Not in this extension.
- [vkGetPhysicalDeviceWaylandPresentationSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceWaylandPresentationSupportKHR.html): Not in this extension.
- [vkGetPhysicalDeviceWin32PresentationSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceWin32PresentationSupportKHR.html): Not in this extension.

## [](#_version_history)Version History

- Revision 1, 2017-02-27 (James Jones)
  
  - Initial draft.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_get_surface_capabilities2)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VK\_AMD\_display\_native\_hdr(3) Manual Page

## Name

VK\_AMD\_display\_native\_hdr - device extension



## [](#_registered_extension_number)Registered Extension Number

214

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)  
and  
[VK\_KHR\_get\_surface\_capabilities2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_surface_capabilities2.html)  
and  
[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html)

## [](#_contact)Contact

- Matthaeus G. Chajdas [\[GitHub\]anteru](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_display_native_hdr%5D%20%40anteru%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_display_native_hdr%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-12-18

**IP Status**

No known IP claims.

**Contributors**

- Matthaeus G. Chajdas, AMD
- Aaron Hagan, AMD
- Aric Cyr, AMD
- Timothy Lottes, AMD
- Derrick Owens, AMD
- Daniel Rakos, AMD

## [](#_description)Description

This extension introduces the following display native HDR features to Vulkan:

- A new [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorSpaceKHR.html) enum for setting the native display color space. For example, this color space would be set by the swapchain to use the native color space in Freesync2 displays.
- Local dimming control

## [](#_new_commands)New Commands

- [vkSetLocalDimmingAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetLocalDimmingAMD.html)

## [](#_new_structures)New Structures

- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html):
  
  - [VkDisplayNativeHdrSurfaceCapabilitiesAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayNativeHdrSurfaceCapabilitiesAMD.html)
- Extending [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html):
  
  - [VkSwapchainDisplayNativeHdrCreateInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainDisplayNativeHdrCreateInfoAMD.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_AMD_DISPLAY_NATIVE_HDR_EXTENSION_NAME`
- `VK_AMD_DISPLAY_NATIVE_HDR_SPEC_VERSION`
- Extending [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorSpaceKHR.html):
  
  - `VK_COLOR_SPACE_DISPLAY_NATIVE_AMD`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DISPLAY_NATIVE_HDR_SURFACE_CAPABILITIES_AMD`
  - `VK_STRUCTURE_TYPE_SWAPCHAIN_DISPLAY_NATIVE_HDR_CREATE_INFO_AMD`

## [](#_issues)Issues

None.

## [](#_examples)Examples

None.

## [](#_version_history)Version History

- Revision 1, 2018-12-18 (Daniel Rakos)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMD_display_native_hdr)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
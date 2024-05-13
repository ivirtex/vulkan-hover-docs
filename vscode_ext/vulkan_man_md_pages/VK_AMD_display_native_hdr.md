# VK_AMD_display_native_hdr(3) Manual Page

## Name

VK_AMD_display_native_hdr - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

214

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Version 1.1](#versions-1.1)  
and  
[VK_KHR_get_surface_capabilities2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_surface_capabilities2.html)  
and  
[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Matthaeus G. Chajdas <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_display_native_hdr%5D%20@anteru%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_display_native_hdr%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>anteru</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension introduces the following display native HDR features to
Vulkan:

- A new [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorSpaceKHR.html) enum for setting the
  native display color space. For example, this color space would be set
  by the swapchain to use the native color space in Freesync2 displays.

- Local dimming control

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkSetLocalDimmingAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetLocalDimmingAMD.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html):

  - [VkDisplayNativeHdrSurfaceCapabilitiesAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayNativeHdrSurfaceCapabilitiesAMD.html)

- Extending [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html):

  - [VkSwapchainDisplayNativeHdrCreateInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainDisplayNativeHdrCreateInfoAMD.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMD_DISPLAY_NATIVE_HDR_EXTENSION_NAME`

- `VK_AMD_DISPLAY_NATIVE_HDR_SPEC_VERSION`

- Extending [VkColorSpaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorSpaceKHR.html):

  - `VK_COLOR_SPACE_DISPLAY_NATIVE_AMD`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DISPLAY_NATIVE_HDR_SURFACE_CAPABILITIES_AMD`

  - `VK_STRUCTURE_TYPE_SWAPCHAIN_DISPLAY_NATIVE_HDR_CREATE_INFO_AMD`

## <a href="#_issues" class="anchor"></a>Issues

None.

## <a href="#_examples" class="anchor"></a>Examples

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-12-18 (Daniel Rakos)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMD_display_native_hdr"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

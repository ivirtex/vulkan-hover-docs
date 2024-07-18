# VK_KHR_surface_protected_capabilities(3) Manual Page

## Name

VK_KHR_surface_protected_capabilities - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

240

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[Version 1.1](#versions-1.1)  
and  
[VK_KHR_get_surface_capabilities2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_surface_capabilities2.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Sandeep Shinde <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_surface_protected_capabilities%5D%20@sashinde%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_surface_protected_capabilities%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>sashinde</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-12-18

**IP Status**  
No known IP claims.

**Contributors**  
- Sandeep Shinde, NVIDIA

- James Jones, NVIDIA

- Daniel Koch, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension extends
[VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html), providing
applications a way to query whether swapchains **can** be created with
the `VK_SWAPCHAIN_CREATE_PROTECTED_BIT_KHR` flag set.

Vulkan 1.1 added (optional) support for protect memory and protected
resources including buffers (`VK_BUFFER_CREATE_PROTECTED_BIT`), images
(`VK_IMAGE_CREATE_PROTECTED_BIT`), and swapchains
(`VK_SWAPCHAIN_CREATE_PROTECTED_BIT_KHR`). However, on implementations
which support multiple windowing systems, not all window systems **may**
be able to provide a protected display path.

This extension provides a way to query if a protected swapchain created
for a surface (and thus a specific windowing system) **can** be
displayed on screen. It extends the existing
[VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html) structure
with a new
[VkSurfaceProtectedCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceProtectedCapabilitiesKHR.html)
structure from which the application **can** obtain information about
support for protected swapchain creation through
[vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html).

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html):

  - [VkSurfaceProtectedCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceProtectedCapabilitiesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SURFACE_PROTECTED_CAPABILITIES_EXTENSION_NAME`

- `VK_KHR_SURFACE_PROTECTED_CAPABILITIES_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_SURFACE_PROTECTED_CAPABILITIES_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-12-18 (Sandeep Shinde, Daniel Koch)

  - Internal revisions.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_surface_protected_capabilities"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

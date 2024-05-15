# VK_KHR_shared_presentable_image(3) Manual Page

## Name

VK_KHR_shared_presentable_image - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

112

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html)  
and  
[VK_KHR_get_surface_capabilities2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_surface_capabilities2.html)  
and  
    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Alon Or-bach <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shared_presentable_image%5D%20@alonorbach%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shared_presentable_image%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>alonorbach</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-03-20

**IP Status**  
No known IP claims.

**Contributors**  
- Alon Or-bach, Samsung Electronics

- Ian Elliott, Google

- Jesse Hall, Google

- Pablo Ceballos, Google

- Chris Forbes, Google

- Jeff Juliano, NVIDIA

- James Jones, NVIDIA

- Daniel Rakos, AMD

- Tobias Hector, Imagination Technologies

- Graham Connor, Imagination Technologies

- Michael Worcester, Imagination Technologies

- Cass Everitt, Oculus

- Johannes Van Waveren, Oculus

## <a href="#_description" class="anchor"></a>Description

This extension extends [`VK_KHR_swapchain`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html) to
enable creation of a shared presentable image. This allows the
application to use the image while the presention engine is accessing
it, in order to reduce the latency between rendering and presentation.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetSwapchainStatusKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSwapchainStatusKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html):

  - [VkSharedPresentSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSharedPresentSurfaceCapabilitiesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SHARED_PRESENTABLE_IMAGE_EXTENSION_NAME`

- `VK_KHR_SHARED_PRESENTABLE_IMAGE_SPEC_VERSION`

- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html):

  - `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`

- Extending [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html):

  - `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`

  - `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_SHARED_PRESENT_SURFACE_CAPABILITIES_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Should we allow a Vulkan WSI swapchain to toggle between normal
usage and shared presentation usage?

**RESOLVED**: No. WSI swapchains are typically recreated with new
properties instead of having their properties changed. This can also
save resources, assuming that fewer images are needed for shared
presentation, and assuming that most VR applications do not need to
switch between normal and shared usage.

2\) Should we have a query for determining how the presentation engine
refresh is triggered?

**RESOLVED**: Yes. This is done via which presentation modes a surface
supports.

3\) Should the object representing a shared presentable image be an
extension of a [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) or a separate
object?

**RESOLVED**: Extension of a swapchain due to overlap in creation
properties and to allow common functionality between shared and normal
presentable images and swapchains.

4\) What should we call the extension and the new structures it creates?

**RESOLVED**: Shared presentable image / shared present.

5\) Should the `minImageCount` and `presentMode` values of the
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) be ignored, or
required to be compatible values?

**RESOLVED**: `minImageCount` must be set to 1, and `presentMode` should
be set to either `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR` or
`VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`.

6\) What should the layout of the shared presentable image be?

**RESOLVED**: After acquiring the shared presentable image, the
application must transition it to the
`VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR` layout prior to it being used.
After this initial transition, any image usage that was requested during
swapchain creation **can** be performed on the image without layout
transitions being performed.

7\) Do we need a new API for the trigger to refresh new content?

**RESOLVED**: [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html) to act as API
to trigger a refresh, as will allow combination with other compatible
extensions to [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html).

8\) How should an application detect a `VK_ERROR_OUT_OF_DATE_KHR` error
on a swapchain using the `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`
present mode?

**RESOLVED**: Introduce
[vkGetSwapchainStatusKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSwapchainStatusKHR.html) to allow
applications to query the status of a swapchain using a shared
presentation mode.

9\) What should subsequent calls to
[vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html) for
`VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR` swapchains be defined to
do?

**RESOLVED**: State that implementations may use it as a hint for
updated content.

10\) Can the ownership of a shared presentable image be transferred to a
different queue?

**RESOLVED**: No. It is not possible to transfer ownership of a shared
presentable image obtained from a swapchain created using
`VK_SHARING_MODE_EXCLUSIVE` after it has been presented.

11\) How should [vkQueueSubmit](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit.html) behave if a command
buffer uses an image from a `VK_ERROR_OUT_OF_DATE_KHR` swapchain?

**RESOLVED**: [vkQueueSubmit](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit.html) is expected to return
the `VK_ERROR_DEVICE_LOST` error.

12\) Can Vulkan provide any guarantee on the order of rendering, to
enable beam chasing?

**RESOLVED**: This could be achieved via use of render passes to ensure
strip rendering.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-03-20 (Alon Or-bach)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_shared_presentable_image"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

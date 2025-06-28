# VK\_KHR\_incremental\_present(3) Manual Page

## Name

VK\_KHR\_incremental\_present - device extension



## [](#_registered_extension_number)Registered Extension Number

85

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html)

## [](#_contact)Contact

- Ian Elliott [\[GitHub\]ianelliottus](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_incremental_present%5D%20%40ianelliottus%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_incremental_present%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-11-02

**IP Status**

No known IP claims.

**Contributors**

- Ian Elliott, Google
- Jesse Hall, Google
- Alon Or-bach, Samsung
- James Jones, NVIDIA
- Daniel Rakos, AMD
- Ray Smith, ARM
- Mika Isojarvi, Google
- Jeff Juliano, NVIDIA
- Jeff Bolz, NVIDIA

## [](#_description)Description

This device extension extends [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html), from the `VK_KHR_swapchain` extension, allowing an application to specify a list of rectangular, modified regions of each image to present. This should be used in situations where an application is only changing a small portion of the presentable images within a swapchain, since it enables the presentation engine to avoid wasting time presenting parts of the surface that have not changed.

This extension is leveraged from the `EGL_KHR_swap_buffers_with_damage` extension.

## [](#_new_structures)New Structures

- [VkPresentRegionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentRegionKHR.html)
- [VkRectLayerKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRectLayerKHR.html)
- Extending [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html):
  
  - [VkPresentRegionsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentRegionsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_INCREMENTAL_PRESENT_EXTENSION_NAME`
- `VK_KHR_INCREMENTAL_PRESENT_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PRESENT_REGIONS_KHR`

## [](#_issues)Issues

1\) How should we handle steroescopic-3D swapchains? We need to add a layer for each rectangle. One approach is to create another structure containing the [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) plus layer, and have [VkPresentRegionsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentRegionsKHR.html) point to an array of that struct. Another approach is to have two parallel arrays, `pRectangles` and `pLayers`, where `pRectangles`\[i] and `pLayers`\[i] must be used together. Which approach should we use, and if the array of a new structure, what should that be called?

**RESOLVED**: Create a new structure, which is a [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) plus a layer, and will be called [VkRectLayerKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRectLayerKHR.html).

2\) Where is the origin of the [VkRectLayerKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRectLayerKHR.html)?

**RESOLVED**: The upper left corner of the presentable image(s) of the swapchain, per the definition of framebuffer coordinates.

3\) Does the rectangular region, [VkRectLayerKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRectLayerKHR.html), specify pixels of the swapchain’s image(s), or of the surface?

**RESOLVED**: Of the image(s). Some presentation engines may scale the pixels of a swapchain’s image(s) to the size of the surface. The size of the swapchain’s image(s) will be consistent, where the size of the surface may vary over time.

4\) What if all of the rectangles for a given swapchain contain a width and/or height of zero?

**RESOLVED**: The application is indicating that no pixels changed since the last present. The presentation engine may use such a hint and not update any pixels for the swapchain. However, all other semantics of [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html) must still be honored, including waiting for semaphores to signal.

5\) When the swapchain is created with `VkSwapchainCreateInfoKHR`::`preTransform` set to a value other than `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, should the rectangular region, [VkRectLayerKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRectLayerKHR.html), be transformed to align with the `preTransform`?

**RESOLVED**: No. The rectangular region in [VkRectLayerKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRectLayerKHR.html) should not be transformed. As such, it may not align with the extents of the swapchain’s image(s). It is the responsibility of the presentation engine to transform the rectangular region. This matches the behavior of the Android presentation engine, which set the precedent.

## [](#_version_history)Version History

- Revision 1, 2016-11-02 (Ian Elliott)
  
  - Internal revisions
- Revision 2, 2021-03-18 (Ian Elliott)
  
  - Clarified alignment of rectangles for presentation engines that support transformed swapchains.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_incremental_present)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
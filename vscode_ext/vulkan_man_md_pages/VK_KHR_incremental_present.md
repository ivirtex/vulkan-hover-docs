# VK_KHR_incremental_present(3) Manual Page

## Name

VK_KHR_incremental_present - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

85

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Ian Elliott <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_incremental_present%5D%20@ianelliottus%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_incremental_present%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>ianelliottus</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This device extension extends
[vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html), from the
[`VK_KHR_swapchain`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html) extension, allowing an
application to specify a list of rectangular, modified regions of each
image to present. This should be used in situations where an application
is only changing a small portion of the presentable images within a
swapchain, since it enables the presentation engine to avoid wasting
time presenting parts of the surface that have not changed.

This extension is leveraged from the `EGL_KHR_swap_buffers_with_damage`
extension.

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkPresentRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentRegionKHR.html)

- [VkRectLayerKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRectLayerKHR.html)

- Extending [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html):

  - [VkPresentRegionsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentRegionsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_INCREMENTAL_PRESENT_EXTENSION_NAME`

- `VK_KHR_INCREMENTAL_PRESENT_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PRESENT_REGIONS_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) How should we handle steroescopic-3D swapchains? We need to add a
layer for each rectangle. One approach is to create another struct
containing the [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) plus layer, and have
[VkPresentRegionsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentRegionsKHR.html) point to an array of
that struct. Another approach is to have two parallel arrays,
`pRectangles` and `pLayers`, where `pRectangles`\[i\] and `pLayers`\[i\]
must be used together. Which approach should we use, and if the array of
a new structure, what should that be called?

**RESOLVED**: Create a new structure, which is a
[VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) plus a layer, and will be called
[VkRectLayerKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRectLayerKHR.html).

2\) Where is the origin of the [VkRectLayerKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRectLayerKHR.html)?

**RESOLVED**: The upper left corner of the presentable image(s) of the
swapchain, per the definition of framebuffer coordinates.

3\) Does the rectangular region, [VkRectLayerKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRectLayerKHR.html),
specify pixels of the swapchain’s image(s), or of the surface?

**RESOLVED**: Of the image(s). Some presentation engines may scale the
pixels of a swapchain’s image(s) to the size of the surface. The size of
the swapchain’s image(s) will be consistent, where the size of the
surface may vary over time.

4\) What if all of the rectangles for a given swapchain contain a width
and/or height of zero?

**RESOLVED**: The application is indicating that no pixels changed since
the last present. The presentation engine may use such a hint and not
update any pixels for the swapchain. However, all other semantics of
[vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html) must still be honored,
including waiting for semaphores to signal.

5\) When the swapchain is created with
`VkSwapchainCreateInfoKHR`::`preTransform` set to a value other than
`VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, should the rectangular region,
[VkRectLayerKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRectLayerKHR.html), be transformed to align with the
`preTransform`?

**RESOLVED**: No. The rectangular region in
[VkRectLayerKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRectLayerKHR.html) should not be transformed. As
such, it may not align with the extents of the swapchain’s image(s). It
is the responsibility of the presentation engine to transform the
rectangular region. This matches the behavior of the Android
presentation engine, which set the precedent.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-11-02 (Ian Elliott)

  - Internal revisions

- Revision 2, 2021-03-18 (Ian Elliott)

  - Clarified alignment of rectangles for presentation engines that
    support transformed swapchains.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_incremental_present"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

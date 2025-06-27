# VkSurfacePresentModeEXT(3) Manual Page

## Name

VkSurfacePresentModeEXT - Structure describing present mode of a surface



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSurfacePresentModeEXT` structure is defined as:

``` c
// Provided by VK_EXT_surface_maintenance1
typedef struct VkSurfacePresentModeEXT {
    VkStructureType     sType;
    void*               pNext;
    VkPresentModeKHR    presentMode;
} VkSurfacePresentModeEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `presentMode` is the presentation mode the swapchain will use.

## <a href="#_description" class="anchor"></a>Description

If the `VkSurfacePresentModeEXT` structure is included in the `pNext`
chain of
[VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html),
the values returned in
[VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html)::`minImageCount`,
[VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html)::`maxImageCount`,
[VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`minScaledImageExtent`,
and
[VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`maxScaledImageExtent`
are valid only for the specified `presentMode`. If `presentMode` is
`VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR` or
`VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`, the per-present mode
image counts **must** both be one. The per-present mode image counts
**may** be less-than or greater-than the image counts returned when
`VkSurfacePresentModeEXT` is not provided.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>If <a
href="VkSwapchainPresentModesCreateInfoEXT.html">VkSwapchainPresentModesCreateInfoEXT</a>
is provided to swapchain creation, the requirements for forward progress
may be less strict. For example, a FIFO swapchain might only require 2
images to guarantee forward progress, but a MAILBOX one might require 4.
Without the per-present image counts, such an implementation would have
to return 4 in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html">VkSurfaceCapabilitiesKHR</a>::<code>minImageCount</code>,
which pessimizes FIFO. Conversely, an implementation may return a low
number for minImageCount, but internally bump the image count when
application queries <a
href="vkGetSwapchainImagesKHR.html">vkGetSwapchainImagesKHR</a>, which
can surprise applications, and is not discoverable until swapchain
creation. Using <code>VkSurfacePresentModeEXT</code> and <a
href="VkSwapchainPresentModesCreateInfoEXT.html">VkSwapchainPresentModesCreateInfoEXT</a>
together effectively removes this problem.</p>
<p><a
href="VkSwapchainPresentModesCreateInfoEXT.html">VkSwapchainPresentModesCreateInfoEXT</a>
is required for the specification to be backwards compatible with
applications that do not know about, or make use of this
feature.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkSurfacePresentModeEXT-presentMode-07780"
  id="VUID-VkSurfacePresentModeEXT-presentMode-07780"></a>
  VUID-VkSurfacePresentModeEXT-presentMode-07780  
  `presentMode` **must** be a value reported by
  [vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html)
  for the specified surface

Valid Usage (Implicit)

- <a href="#VUID-VkSurfacePresentModeEXT-sType-sType"
  id="VUID-VkSurfacePresentModeEXT-sType-sType"></a>
  VUID-VkSurfacePresentModeEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_PRESENT_MODE_EXT`

- <a href="#VUID-VkSurfacePresentModeEXT-presentMode-parameter"
  id="VUID-VkSurfacePresentModeEXT-presentMode-parameter"></a>
  VUID-VkSurfacePresentModeEXT-presentMode-parameter  
  `presentMode` **must** be a valid
  [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_surface_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_surface_maintenance1.html),
[VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSurfacePresentModeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

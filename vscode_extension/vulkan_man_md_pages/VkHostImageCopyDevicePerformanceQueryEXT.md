# VkHostImageCopyDevicePerformanceQueryEXT(3) Manual Page

## Name

VkHostImageCopyDevicePerformanceQueryEXT - Struct containing information
about optimality of device access



## <a href="#_c_specification" class="anchor"></a>C Specification

To query if using `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` has a negative
impact on device performance when accessing an image, add
`VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` to
[VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`usage`,
and add a `VkHostImageCopyDevicePerformanceQueryEXT` structure to the
`pNext` chain of a
[VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html) structure
passed to
[vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html).
This structure is defined as:

``` c
// Provided by VK_EXT_host_image_copy
typedef struct VkHostImageCopyDevicePerformanceQueryEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           optimalDeviceAccess;
    VkBool32           identicalMemoryLayout;
} VkHostImageCopyDevicePerformanceQueryEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `optimalDeviceAccess` returns `VK_TRUE` if use of host image copy has
  no adverse effect on device access performance, compared to an image
  that is created with exact same creation parameters, and bound to the
  same [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html), except that
  `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` is replaced with
  `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` and
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT`.

- `identicalMemoryLayout` returns `VK_TRUE` if use of host image copy
  has no impact on memory layout compared to an image that is created
  with exact same creation parameters, and bound to the same
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html), except that
  `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` is replaced with
  `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` and
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT`.

## <a href="#_description" class="anchor"></a>Description

The implementation **may** return `VK_FALSE` in `optimalDeviceAccess` if
`identicalMemoryLayout` is `VK_FALSE`. If `identicalMemoryLayout` is
`VK_TRUE`, `optimalDeviceAccess` **must** be `VK_TRUE`.

The implementation **may** return `VK_TRUE` in `optimalDeviceAccess`
while `identicalMemoryLayout` is `VK_FALSE`. In this situation, any
device performance impact **should** not be measurable.

If
[VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`format`
is a block-compressed format and
[vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
returns `VK_SUCCESS`, the implementation **must** return `VK_TRUE` in
`optimalDeviceAccess`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Applications can make use of <code>optimalDeviceAccess</code> to
determine their resource copying strategy. If a resource is expected to
be accessed more on device than on the host, and the implementation
considers the resource sub-optimally accessed, it is likely better to
use device copies instead.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Layout not being identical yet still considered optimal for device
access could happen if the implementation has different memory layout
patterns, some of which are easier to access on the host.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The most practical reason for <code>optimalDeviceAccess</code> to be
<code>VK_FALSE</code> is that host image access may disable framebuffer
compression where it would otherwise have been enabled. This represents
far more efficient host image access since no compression algorithm is
required to read or write to the image, but it would impact device
access performance. Some implementations may only set
<code>optimalDeviceAccess</code> to <code>VK_FALSE</code> if certain
conditions are met, such as specific image usage flags or creation
flags.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkHostImageCopyDevicePerformanceQueryEXT-sType-sType"
  id="VUID-VkHostImageCopyDevicePerformanceQueryEXT-sType-sType"></a>
  VUID-VkHostImageCopyDevicePerformanceQueryEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_HOST_IMAGE_COPY_DEVICE_PERFORMANCE_QUERY_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_host_image_copy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_image_copy.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkHostImageCopyDevicePerformanceQueryEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

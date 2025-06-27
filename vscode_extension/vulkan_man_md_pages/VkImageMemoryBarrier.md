# VkImageMemoryBarrier(3) Manual Page

## Name

VkImageMemoryBarrier - Structure specifying the parameters of an image
memory barrier



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageMemoryBarrier` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkImageMemoryBarrier {
    VkStructureType            sType;
    const void*                pNext;
    VkAccessFlags              srcAccessMask;
    VkAccessFlags              dstAccessMask;
    VkImageLayout              oldLayout;
    VkImageLayout              newLayout;
    uint32_t                   srcQueueFamilyIndex;
    uint32_t                   dstQueueFamilyIndex;
    VkImage                    image;
    VkImageSubresourceRange    subresourceRange;
} VkImageMemoryBarrier;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `srcAccessMask` is a bitmask of
  [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html) specifying a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
  target="_blank" rel="noopener">source access mask</a>.

- `dstAccessMask` is a bitmask of
  [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html) specifying a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
  target="_blank" rel="noopener">destination access mask</a>.

- `oldLayout` is the old layout in an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
  target="_blank" rel="noopener">image layout transition</a>.

- `newLayout` is the new layout in an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
  target="_blank" rel="noopener">image layout transition</a>.

- `srcQueueFamilyIndex` is the source queue family for a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
  target="_blank" rel="noopener">queue family ownership transfer</a>.

- `dstQueueFamilyIndex` is the destination queue family for a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
  target="_blank" rel="noopener">queue family ownership transfer</a>.

- `image` is a handle to the image affected by this barrier.

- `subresourceRange` describes the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views"
  target="_blank" rel="noopener">image subresource range</a> within
  `image` that is affected by this barrier.

## <a href="#_description" class="anchor"></a>Description

The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> is limited to access to
memory through the specified image subresource range, via access types
in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
target="_blank" rel="noopener">source access mask</a> specified by
`srcAccessMask`. If `srcAccessMask` includes `VK_ACCESS_HOST_WRITE_BIT`,
memory writes performed by that access type are also made visible, as
that access type is not performed through a resource.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> is limited to access to
memory through the specified image subresource range, via access types
in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
target="_blank" rel="noopener">destination access mask</a> specified by
`dstAccessMask`. If `dstAccessMask` includes `VK_ACCESS_HOST_WRITE_BIT`
or `VK_ACCESS_HOST_READ_BIT`, available memory writes are also made
visible to accesses of those types, as those access types are not
performed through a resource.

If `srcQueueFamilyIndex` is not equal to `dstQueueFamilyIndex`, and
`srcQueueFamilyIndex` is equal to the current queue family, then the
memory barrier defines a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers-release"
target="_blank" rel="noopener">queue family release operation</a> for
the specified image subresource range, and the second synchronization
scope of the calling command does not apply to this operation.

If `dstQueueFamilyIndex` is not equal to `srcQueueFamilyIndex`, and
`dstQueueFamilyIndex` is equal to the current queue family, then the
memory barrier defines a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers-acquire"
target="_blank" rel="noopener">queue family acquire operation</a> for
the specified image subresource range, and the first synchronization
scope of the calling command does not apply to this operation.

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-synchronization2"
target="_blank" rel="noopener"><code>synchronization2</code></a> feature
is not enabled or `oldLayout` is not equal to `newLayout`, `oldLayout`
and `newLayout` define an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
target="_blank" rel="noopener">image layout transition</a> for the
specified image subresource range.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-synchronization2"
target="_blank" rel="noopener"><code>synchronization2</code></a> feature
is enabled, when the old and new layout are equal, the layout values are
ignored - data is preserved no matter what values are specified, or what
layout the image is currently in.</p></td>
</tr>
</tbody>
</table>

If `image` has a multi-planar format and the image is *disjoint*, then
including `VK_IMAGE_ASPECT_COLOR_BIT` in the `aspectMask` member of
`subresourceRange` is equivalent to including
`VK_IMAGE_ASPECT_PLANE_0_BIT`, `VK_IMAGE_ASPECT_PLANE_1_BIT`, and (for
three-plane formats only) `VK_IMAGE_ASPECT_PLANE_2_BIT`.

Valid Usage

- <a href="#VUID-VkImageMemoryBarrier-oldLayout-01208"
  id="VUID-VkImageMemoryBarrier-oldLayout-01208"></a>
  VUID-VkImageMemoryBarrier-oldLayout-01208  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL` then `image` **must** have
  been created with `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier-oldLayout-01209"
  id="VUID-VkImageMemoryBarrier-oldLayout-01209"></a>
  VUID-VkImageMemoryBarrier-oldLayout-01209  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` then `image`
  **must** have been created with
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier-oldLayout-01210"
  id="VUID-VkImageMemoryBarrier-oldLayout-01210"></a>
  VUID-VkImageMemoryBarrier-oldLayout-01210  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL` then `image`
  **must** have been created with
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier-oldLayout-01211"
  id="VUID-VkImageMemoryBarrier-oldLayout-01211"></a>
  VUID-VkImageMemoryBarrier-oldLayout-01211  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL` then `image` **must** have
  been created with `VK_IMAGE_USAGE_SAMPLED_BIT` or
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier-oldLayout-01212"
  id="VUID-VkImageMemoryBarrier-oldLayout-01212"></a>
  VUID-VkImageMemoryBarrier-oldLayout-01212  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_SRC_BIT`

- <a href="#VUID-VkImageMemoryBarrier-oldLayout-01213"
  id="VUID-VkImageMemoryBarrier-oldLayout-01213"></a>
  VUID-VkImageMemoryBarrier-oldLayout-01213  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT`

- <a href="#VUID-VkImageMemoryBarrier-oldLayout-01197"
  id="VUID-VkImageMemoryBarrier-oldLayout-01197"></a>
  VUID-VkImageMemoryBarrier-oldLayout-01197  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), `oldLayout`
  **must** be `VK_IMAGE_LAYOUT_UNDEFINED` or the current layout of the
  image subresources affected by the barrier

- <a href="#VUID-VkImageMemoryBarrier-newLayout-01198"
  id="VUID-VkImageMemoryBarrier-newLayout-01198"></a>
  VUID-VkImageMemoryBarrier-newLayout-01198  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), `newLayout`
  **must** not be `VK_IMAGE_LAYOUT_UNDEFINED` or
  `VK_IMAGE_LAYOUT_PREINITIALIZED`

- <a href="#VUID-VkImageMemoryBarrier-oldLayout-01658"
  id="VUID-VkImageMemoryBarrier-oldLayout-01658"></a>
  VUID-VkImageMemoryBarrier-oldLayout-01658  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL` then
  `image` **must** have been created with
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier-oldLayout-01659"
  id="VUID-VkImageMemoryBarrier-oldLayout-01659"></a>
  VUID-VkImageMemoryBarrier-oldLayout-01659  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` then
  `image` **must** have been created with
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04065"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04065"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04065  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL` then `image` **must** have
  been created with at least one of
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`,
  `VK_IMAGE_USAGE_SAMPLED_BIT`, or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04066"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04066"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04066  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` then `image` **must** have
  been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` set

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04067"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04067"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04067  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL` then `image` **must** have
  been created with at least one of
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`,
  `VK_IMAGE_USAGE_SAMPLED_BIT`, or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04068"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04068"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04068  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` then `image` **must**
  have been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
  set

- <a href="#VUID-VkImageMemoryBarrier-synchronization2-07793"
  id="VUID-VkImageMemoryBarrier-synchronization2-07793"></a>
  VUID-VkImageMemoryBarrier-synchronization2-07793  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `oldLayout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or
  `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`

- <a href="#VUID-VkImageMemoryBarrier-synchronization2-07794"
  id="VUID-VkImageMemoryBarrier-synchronization2-07794"></a>
  VUID-VkImageMemoryBarrier-synchronization2-07794  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `newLayout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or
  `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-03938"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-03938"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-03938  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL`,
  `image` **must** have been created with
  `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-03939"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-03939"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-03939  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL`,
  `image` **must** have been created with at least one of
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`,
  `VK_IMAGE_USAGE_SAMPLED_BIT`, or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier-oldLayout-02088"
  id="VUID-VkImageMemoryBarrier-oldLayout-02088"></a>
  VUID-VkImageMemoryBarrier-oldLayout-02088  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR` then
  `image` **must** have been created with
  `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR` set

- <a href="#VUID-VkImageMemoryBarrier-image-09117"
  id="VUID-VkImageMemoryBarrier-image-09117"></a>
  VUID-VkImageMemoryBarrier-image-09117  
  If `image` was created with a sharing mode of
  `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` and
  `dstQueueFamilyIndex` are not equal, `srcQueueFamilyIndex` **must** be
  `VK_QUEUE_FAMILY_EXTERNAL`, `VK_QUEUE_FAMILY_FOREIGN_EXT`, or a valid
  queue family

- <a href="#VUID-VkImageMemoryBarrier-image-09118"
  id="VUID-VkImageMemoryBarrier-image-09118"></a>
  VUID-VkImageMemoryBarrier-image-09118  
  If `image` was created with a sharing mode of
  `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` and
  `dstQueueFamilyIndex` are not equal, `dstQueueFamilyIndex` **must** be
  `VK_QUEUE_FAMILY_EXTERNAL`, `VK_QUEUE_FAMILY_FOREIGN_EXT`, or a valid
  queue family

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04070"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04070"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04070  
  If `srcQueueFamilyIndex` is not equal to `dstQueueFamilyIndex`, at
  least one of `srcQueueFamilyIndex` or `dstQueueFamilyIndex` **must**
  not be `VK_QUEUE_FAMILY_EXTERNAL` or `VK_QUEUE_FAMILY_FOREIGN_EXT`

- <a href="#VUID-VkImageMemoryBarrier-None-09119"
  id="VUID-VkImageMemoryBarrier-None-09119"></a>
  VUID-VkImageMemoryBarrier-None-09119  
  If the [VK_KHR_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html) extension
  is not enabled, and the value of
  [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion` used to
  create the [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) is not greater than or equal
  to Version 1.1, `srcQueueFamilyIndex` **must** not be
  `VK_QUEUE_FAMILY_EXTERNAL`

- <a href="#VUID-VkImageMemoryBarrier-None-09120"
  id="VUID-VkImageMemoryBarrier-None-09120"></a>
  VUID-VkImageMemoryBarrier-None-09120  
  If the [VK_KHR_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html) extension
  is not enabled, and the value of
  [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion` used to
  create the [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) is not greater than or equal
  to Version 1.1, `dstQueueFamilyIndex` **must** not be
  `VK_QUEUE_FAMILY_EXTERNAL`

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-09121"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-09121"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-09121  
  If the [VK_EXT_queue_family_foreign](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_queue_family_foreign.html)
  extension is not enabled `srcQueueFamilyIndex` **must** not be
  `VK_QUEUE_FAMILY_FOREIGN_EXT`

- <a href="#VUID-VkImageMemoryBarrier-dstQueueFamilyIndex-09122"
  id="VUID-VkImageMemoryBarrier-dstQueueFamilyIndex-09122"></a>
  VUID-VkImageMemoryBarrier-dstQueueFamilyIndex-09122  
  If the [VK_EXT_queue_family_foreign](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_queue_family_foreign.html)
  extension is not enabled `dstQueueFamilyIndex` **must** not be
  `VK_QUEUE_FAMILY_FOREIGN_EXT`

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07120"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07120"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07120  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_DECODE_SRC_KHR`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_VIDEO_DECODE_SRC_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07121"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07121"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07121  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_DECODE_DST_KHR`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07122"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07122"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07122  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_DECODE_DPB_KHR`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07123"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07123"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07123  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_ENCODE_SRC_KHR`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07124"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07124"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07124  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DST_KHR`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_VIDEO_ENCODE_DST_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07125"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07125"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07125  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DPB_KHR`
  then `image` **must** have been created with
  `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07006"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07006"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07006  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` then `image`
  **must** have been created with either the
  `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` usage bits, and the
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_SAMPLED_BIT`
  usage bits, and the `VK_IMAGE_USAGE_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`
  usage bit

- <a href="#VUID-VkImageMemoryBarrier-attachmentFeedbackLoopLayout-07313"
  id="VUID-VkImageMemoryBarrier-attachmentFeedbackLoopLayout-07313"></a>
  VUID-VkImageMemoryBarrier-attachmentFeedbackLoopLayout-07313  
  If the
  [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout)
  feature is not enabled, `newLayout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`

- <a href="#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-09550"
  id="VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-09550"></a>
  VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-09550  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue
  family ownership transfer](#synchronization-queue-transfers) or
  `oldLayout` and `newLayout` define an [image layout
  transition](#synchronization-image-layout-transitions), and
  `oldLayout` or `newLayout` is
  `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR` then `image` **must** have
  been created with either `VK_IMAGE_USAGE_STORAGE_BIT`, or with both
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` and either of
  `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkImageMemoryBarrier-dynamicRenderingLocalRead-09551"
  id="VUID-VkImageMemoryBarrier-dynamicRenderingLocalRead-09551"></a>
  VUID-VkImageMemoryBarrier-dynamicRenderingLocalRead-09551  
  If the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, `oldLayout` **must** not be
  `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR`

- <a href="#VUID-VkImageMemoryBarrier-dynamicRenderingLocalRead-09552"
  id="VUID-VkImageMemoryBarrier-dynamicRenderingLocalRead-09552"></a>
  VUID-VkImageMemoryBarrier-dynamicRenderingLocalRead-09552  
  If the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, `newLayout` **must** not be
  `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR`

<!-- -->

- <a href="#VUID-VkImageMemoryBarrier-subresourceRange-01486"
  id="VUID-VkImageMemoryBarrier-subresourceRange-01486"></a>
  VUID-VkImageMemoryBarrier-subresourceRange-01486  
  `subresourceRange.baseMipLevel` **must** be less than the `mipLevels`
  specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image`
  was created

- <a href="#VUID-VkImageMemoryBarrier-subresourceRange-01724"
  id="VUID-VkImageMemoryBarrier-subresourceRange-01724"></a>
  VUID-VkImageMemoryBarrier-subresourceRange-01724  
  If `subresourceRange.levelCount` is not `VK_REMAINING_MIP_LEVELS`,
  `subresourceRange.baseMipLevel` + `subresourceRange.levelCount`
  **must** be less than or equal to the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-VkImageMemoryBarrier-subresourceRange-01488"
  id="VUID-VkImageMemoryBarrier-subresourceRange-01488"></a>
  VUID-VkImageMemoryBarrier-subresourceRange-01488  
  `subresourceRange.baseArrayLayer` **must** be less than the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `image` was created

- <a href="#VUID-VkImageMemoryBarrier-subresourceRange-01725"
  id="VUID-VkImageMemoryBarrier-subresourceRange-01725"></a>
  VUID-VkImageMemoryBarrier-subresourceRange-01725  
  If `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `subresourceRange.baseArrayLayer` + `subresourceRange.layerCount`
  **must** be less than or equal to the `arrayLayers` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-VkImageMemoryBarrier-image-01932"
  id="VUID-VkImageMemoryBarrier-image-01932"></a>
  VUID-VkImageMemoryBarrier-image-01932  
  If `image` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkImageMemoryBarrier-image-09241"
  id="VUID-VkImageMemoryBarrier-image-09241"></a>
  VUID-VkImageMemoryBarrier-image-09241  
  If `image` has a color format that is single-plane, then the
  `aspectMask` member of `subresourceRange` **must** be
  `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-VkImageMemoryBarrier-image-09242"
  id="VUID-VkImageMemoryBarrier-image-09242"></a>
  VUID-VkImageMemoryBarrier-image-09242  
  If `image` has a color format and is not *disjoint*, then the
  `aspectMask` member of `subresourceRange` **must** be
  `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-VkImageMemoryBarrier-image-01672"
  id="VUID-VkImageMemoryBarrier-image-01672"></a>
  VUID-VkImageMemoryBarrier-image-01672  
  If `image` has a multi-planar format and the image is *disjoint*, then
  the `aspectMask` member of `subresourceRange` **must** include at
  least one [multi-planar aspect mask](#formats-planes-image-aspect) bit
  or `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-VkImageMemoryBarrier-image-03320"
  id="VUID-VkImageMemoryBarrier-image-03320"></a>
  VUID-VkImageMemoryBarrier-image-03320  
  If `image` has a depth/stencil format with both depth and stencil and
  the
  [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts)
  feature is not enabled, then the `aspectMask` member of
  `subresourceRange` **must** include both `VK_IMAGE_ASPECT_DEPTH_BIT`
  and `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-VkImageMemoryBarrier-image-03319"
  id="VUID-VkImageMemoryBarrier-image-03319"></a>
  VUID-VkImageMemoryBarrier-image-03319  
  If `image` has a depth/stencil format with both depth and stencil and
  the
  [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts)
  feature is enabled, then the `aspectMask` member of `subresourceRange`
  **must** include either or both `VK_IMAGE_ASPECT_DEPTH_BIT` and
  `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-VkImageMemoryBarrier-aspectMask-08702"
  id="VUID-VkImageMemoryBarrier-aspectMask-08702"></a>
  VUID-VkImageMemoryBarrier-aspectMask-08702  
  If the `aspectMask` member of `subresourceRange` includes
  `VK_IMAGE_ASPECT_DEPTH_BIT`, `oldLayout` and `newLayout` **must** not
  be one of `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkImageMemoryBarrier-aspectMask-08703"
  id="VUID-VkImageMemoryBarrier-aspectMask-08703"></a>
  VUID-VkImageMemoryBarrier-aspectMask-08703  
  If the `aspectMask` member of `subresourceRange` includes
  `VK_IMAGE_ASPECT_STENCIL_BIT`, `oldLayout` and `newLayout` **must**
  not be one of `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkImageMemoryBarrier-subresourceRange-09601"
  id="VUID-VkImageMemoryBarrier-subresourceRange-09601"></a>
  VUID-VkImageMemoryBarrier-subresourceRange-09601  
  `subresourceRange.aspectMask` **must** be valid for the `format` the
  `image` was created with

- <a href="#VUID-VkImageMemoryBarrier-None-09052"
  id="VUID-VkImageMemoryBarrier-None-09052"></a>
  VUID-VkImageMemoryBarrier-None-09052  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-synchronization2"
  target="_blank" rel="noopener"><code>synchronization2</code></a>
  feature is not enabled, and `image` was created with a sharing mode of
  `VK_SHARING_MODE_CONCURRENT`, at least one of `srcQueueFamilyIndex`
  and `dstQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_IGNORED`

- <a href="#VUID-VkImageMemoryBarrier-None-09053"
  id="VUID-VkImageMemoryBarrier-None-09053"></a>
  VUID-VkImageMemoryBarrier-None-09053  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-synchronization2"
  target="_blank" rel="noopener"><code>synchronization2</code></a>
  feature is not enabled, and `image` was created with a sharing mode of
  `VK_SHARING_MODE_CONCURRENT`, `srcQueueFamilyIndex` **must** be
  `VK_QUEUE_FAMILY_IGNORED` or `VK_QUEUE_FAMILY_EXTERNAL`

- <a href="#VUID-VkImageMemoryBarrier-None-09054"
  id="VUID-VkImageMemoryBarrier-None-09054"></a>
  VUID-VkImageMemoryBarrier-None-09054  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-synchronization2"
  target="_blank" rel="noopener"><code>synchronization2</code></a>
  feature is not enabled, and `image` was created with a sharing mode of
  `VK_SHARING_MODE_CONCURRENT`, `dstQueueFamilyIndex` **must** be
  `VK_QUEUE_FAMILY_IGNORED` or `VK_QUEUE_FAMILY_EXTERNAL`

Valid Usage (Implicit)

- <a href="#VUID-VkImageMemoryBarrier-sType-sType"
  id="VUID-VkImageMemoryBarrier-sType-sType"></a>
  VUID-VkImageMemoryBarrier-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER`

- <a href="#VUID-VkImageMemoryBarrier-pNext-pNext"
  id="VUID-VkImageMemoryBarrier-pNext-pNext"></a>
  VUID-VkImageMemoryBarrier-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkExternalMemoryAcquireUnmodifiedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryAcquireUnmodifiedEXT.html)
  or [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationsInfoEXT.html)

- <a href="#VUID-VkImageMemoryBarrier-sType-unique"
  id="VUID-VkImageMemoryBarrier-sType-unique"></a>
  VUID-VkImageMemoryBarrier-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkImageMemoryBarrier-oldLayout-parameter"
  id="VUID-VkImageMemoryBarrier-oldLayout-parameter"></a>
  VUID-VkImageMemoryBarrier-oldLayout-parameter  
  `oldLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)
  value

- <a href="#VUID-VkImageMemoryBarrier-newLayout-parameter"
  id="VUID-VkImageMemoryBarrier-newLayout-parameter"></a>
  VUID-VkImageMemoryBarrier-newLayout-parameter  
  `newLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)
  value

- <a href="#VUID-VkImageMemoryBarrier-image-parameter"
  id="VUID-VkImageMemoryBarrier-image-parameter"></a>
  VUID-VkImageMemoryBarrier-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkImageMemoryBarrier-subresourceRange-parameter"
  id="VUID-VkImageMemoryBarrier-subresourceRange-parameter"></a>
  VUID-VkImageMemoryBarrier-subresourceRange-parameter  
  `subresourceRange` **must** be a valid
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAccessFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier.html),
[vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageMemoryBarrier"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

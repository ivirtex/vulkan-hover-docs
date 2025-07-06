# VkImageMemoryBarrier(3) Manual Page

## Name

VkImageMemoryBarrier - Structure specifying the parameters of an image memory barrier



## [](#_c_specification)C Specification

The `VkImageMemoryBarrier` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcAccessMask` is a bitmask of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) specifying a [source access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks).
- `dstAccessMask` is a bitmask of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) specifying a [destination access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks).
- `oldLayout` is the old layout in an [image layout transition](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-image-layout-transitions).
- `newLayout` is the new layout in an [image layout transition](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-image-layout-transitions).
- `srcQueueFamilyIndex` is the source queue family for a [queue family ownership transfer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers).
- `dstQueueFamilyIndex` is the destination queue family for a [queue family ownership transfer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers).
- `image` is a handle to the image affected by this barrier.
- `subresourceRange` describes the [image subresource range](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views) within `image` that is affected by this barrier.

## [](#_description)Description

The first [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) is limited to access to memory through the specified image subresource range, via access types in the [source access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks) specified by `srcAccessMask`. If `srcAccessMask` includes `VK_ACCESS_HOST_WRITE_BIT`, memory writes performed by that access type are also made visible, as that access type is not performed through a resource.

The second [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) is limited to access to memory through the specified image subresource range, via access types in the [destination access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks) specified by `dstAccessMask`. If `dstAccessMask` includes `VK_ACCESS_HOST_WRITE_BIT` or `VK_ACCESS_HOST_READ_BIT`, available memory writes are also made visible to accesses of those types, as those access types are not performed through a resource.

If `srcQueueFamilyIndex` is not equal to `dstQueueFamilyIndex`, and `srcQueueFamilyIndex` is equal to the current queue family, then the memory barrier defines a [queue family release operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-release) for the specified image subresource range, and if `dependencyFlags` did not include `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR`, the second synchronization scope of the calling command does not apply to this operation.

If `dstQueueFamilyIndex` is not equal to `srcQueueFamilyIndex`, and `dstQueueFamilyIndex` is equal to the current queue family, then the memory barrier defines a [queue family acquire operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-acquire) for the specified image subresource range, and if `dependencyFlags` did not include `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR`, the first synchronization scope of the calling command does not apply to this operation.

If the [`synchronization2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-synchronization2) feature is not enabled or `oldLayout` is not equal to `newLayout`, `oldLayout` and `newLayout` define an [image layout transition](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-image-layout-transitions) for the specified image subresource range.

Note

If the [`synchronization2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-synchronization2) feature is enabled, when the old and new layout are equal, the layout values are ignored - data is preserved no matter what values are specified, or what layout the image is currently in.

If `image` is a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` and the [`maintenance9`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance9) feature is enabled, the `baseArrayLayer` and `layerCount` members of `subresourceRange` specify the subset of slices of the 3D image affected by the memory barrier, including the layout transition. Any slices of a 3D image not included in `subresourceRange` are not affected by the memory barrier and remain in their existing layout.

If `image` has a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar) and the image is *disjoint*, then including `VK_IMAGE_ASPECT_COLOR_BIT` in the `aspectMask` member of `subresourceRange` is equivalent to including `VK_IMAGE_ASPECT_PLANE_0_BIT`, `VK_IMAGE_ASPECT_PLANE_1_BIT`, and (for three-plane formats only) `VK_IMAGE_ASPECT_PLANE_2_BIT`.

Valid Usage

- [](#VUID-VkImageMemoryBarrier-oldLayout-01208)VUID-VkImageMemoryBarrier-oldLayout-01208  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier-oldLayout-01209)VUID-VkImageMemoryBarrier-oldLayout-01209  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier-oldLayout-01210)VUID-VkImageMemoryBarrier-oldLayout-01210  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier-oldLayout-01211)VUID-VkImageMemoryBarrier-oldLayout-01211  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_SAMPLED_BIT` or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier-oldLayout-01212)VUID-VkImageMemoryBarrier-oldLayout-01212  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_TRANSFER_SRC_BIT`
- [](#VUID-VkImageMemoryBarrier-oldLayout-01213)VUID-VkImageMemoryBarrier-oldLayout-01213  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_TRANSFER_DST_BIT`
- [](#VUID-VkImageMemoryBarrier-oldLayout-01197)VUID-VkImageMemoryBarrier-oldLayout-01197  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), `oldLayout` **must** be `VK_IMAGE_LAYOUT_UNDEFINED` or the current layout of the image subresources affected by the barrier
- [](#VUID-VkImageMemoryBarrier-oldLayout-10767)VUID-VkImageMemoryBarrier-oldLayout-10767  
  If the [zeroInitializeDeviceMemory](#features-zeroInitializeDeviceMemory) feature is not enabled, `oldLayout` **must** not be `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT`
- [](#VUID-VkImageMemoryBarrier-oldLayout-10768)VUID-VkImageMemoryBarrier-oldLayout-10768  
  If `oldLayout` is `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT`, then all subresources **must** be included in the barrier
- [](#VUID-VkImageMemoryBarrier-newLayout-01198)VUID-VkImageMemoryBarrier-newLayout-01198  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), `newLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED` or `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT` or `VK_IMAGE_LAYOUT_PREINITIALIZED`
- [](#VUID-VkImageMemoryBarrier-oldLayout-01658)VUID-VkImageMemoryBarrier-oldLayout-01658  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier-oldLayout-01659)VUID-VkImageMemoryBarrier-oldLayout-01659  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04065)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04065  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL` then `image` **must** have been created with at least one of `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`, `VK_IMAGE_USAGE_SAMPLED_BIT`, or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04066)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04066  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` set
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04067)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04067  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL` then `image` **must** have been created with at least one of `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`, `VK_IMAGE_USAGE_SAMPLED_BIT`, or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04068)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-04068  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` then `image` **must** have been created with `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` set
- [](#VUID-VkImageMemoryBarrier-synchronization2-07793)VUID-VkImageMemoryBarrier-synchronization2-07793  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `oldLayout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- [](#VUID-VkImageMemoryBarrier-synchronization2-07794)VUID-VkImageMemoryBarrier-synchronization2-07794  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `newLayout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-03938)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-03938  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL`, `image` **must** have been created with `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-03939)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-03939  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL`, `image` **must** have been created with at least one of `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`, `VK_IMAGE_USAGE_SAMPLED_BIT`, or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier-oldLayout-02088)VUID-VkImageMemoryBarrier-oldLayout-02088  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR` set
- [](#VUID-VkImageMemoryBarrier-image-09117)VUID-VkImageMemoryBarrier-image-09117  
  If `image` was created with a sharing mode of `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` and `dstQueueFamilyIndex` are not equal, `srcQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_EXTERNAL`, `VK_QUEUE_FAMILY_FOREIGN_EXT`, or a valid queue family
- [](#VUID-VkImageMemoryBarrier-image-09118)VUID-VkImageMemoryBarrier-image-09118  
  If `image` was created with a sharing mode of `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` and `dstQueueFamilyIndex` are not equal, `dstQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_EXTERNAL`, `VK_QUEUE_FAMILY_FOREIGN_EXT`, or a valid queue family
- [](#VUID-VkImageMemoryBarrier-None-09119)VUID-VkImageMemoryBarrier-None-09119  
  If the [VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html) extension is not enabled, and the value of [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion` used to create the [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) is not greater than or equal to Version 1.1, `srcQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_EXTERNAL`
- [](#VUID-VkImageMemoryBarrier-None-09120)VUID-VkImageMemoryBarrier-None-09120  
  If the [VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html) extension is not enabled, and the value of [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion` used to create the [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) is not greater than or equal to Version 1.1, `dstQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_EXTERNAL`
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-09121)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-09121  
  If the [VK\_EXT\_queue\_family\_foreign](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_queue_family_foreign.html) extension is not enabled `srcQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_FOREIGN_EXT`
- [](#VUID-VkImageMemoryBarrier-dstQueueFamilyIndex-09122)VUID-VkImageMemoryBarrier-dstQueueFamilyIndex-09122  
  If the [VK\_EXT\_queue\_family\_foreign](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_queue_family_foreign.html) extension is not enabled `dstQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_FOREIGN_EXT`
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07120)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07120  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_DECODE_SRC_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_VIDEO_DECODE_SRC_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07121)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07121  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_DECODE_DST_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07122)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07122  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_DECODE_DPB_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07123)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07123  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_ENCODE_SRC_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07124)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07124  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DST_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_VIDEO_ENCODE_DST_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07125)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07125  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DPB_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-10287)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-10287  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_VIDEO_ENCODE_QUANTIZATION_MAP_KHR` then `image` **must** have been created with `VK_IMAGE_USAGE_VIDEO_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR` or `VK_IMAGE_USAGE_VIDEO_ENCODE_EMPHASIS_MAP_BIT_KHR`
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07006)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-07006  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` then `image` **must** have been created with either the `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` usage bits, and the `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_SAMPLED_BIT` usage bits, and the `VK_IMAGE_USAGE_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT` usage bit
- [](#VUID-VkImageMemoryBarrier-attachmentFeedbackLoopLayout-07313)VUID-VkImageMemoryBarrier-attachmentFeedbackLoopLayout-07313  
  If the [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout) feature is not enabled, `newLayout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`
- [](#VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-09550)VUID-VkImageMemoryBarrier-srcQueueFamilyIndex-09550  
  If `srcQueueFamilyIndex` and `dstQueueFamilyIndex` define a [queue family ownership transfer](#synchronization-queue-transfers) or `oldLayout` and `newLayout` define an [image layout transition](#synchronization-image-layout-transitions), and `oldLayout` or `newLayout` is `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ` then `image` **must** have been created with either `VK_IMAGE_USAGE_STORAGE_BIT`, or with both `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` and either of `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkImageMemoryBarrier-dynamicRenderingLocalRead-09551)VUID-VkImageMemoryBarrier-dynamicRenderingLocalRead-09551  
  If the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, `oldLayout` **must** not be `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ`
- [](#VUID-VkImageMemoryBarrier-dynamicRenderingLocalRead-09552)VUID-VkImageMemoryBarrier-dynamicRenderingLocalRead-09552  
  If the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, `newLayout` **must** not be `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ`

<!--THE END-->

- [](#VUID-VkImageMemoryBarrier-subresourceRange-01486)VUID-VkImageMemoryBarrier-subresourceRange-01486  
  `subresourceRange.baseMipLevel` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkImageMemoryBarrier-subresourceRange-01724)VUID-VkImageMemoryBarrier-subresourceRange-01724  
  If `subresourceRange.levelCount` is not `VK_REMAINING_MIP_LEVELS`, `subresourceRange.baseMipLevel` + `subresourceRange.levelCount` **must** be less than or equal to the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkImageMemoryBarrier-subresourceRange-01488)VUID-VkImageMemoryBarrier-subresourceRange-01488  
  If `image` is not a 3D image or was created without `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, or the [`maintenance9`](#features-maintenance9) feature is not enabled, `subresourceRange.baseArrayLayer` **must** be less than the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkImageMemoryBarrier-maintenance9-10798)VUID-VkImageMemoryBarrier-maintenance9-10798  
  If the [`maintenance9`](#features-maintenance9) feature is enabled and `image` is a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, `subresourceRange.baseArrayLayer` **must** be less than the depth computed from `baseMipLevel` and `extent.depth` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created, according to the formula defined in [Image Mip Level Sizing](#resources-image-mip-level-sizing)
- [](#VUID-VkImageMemoryBarrier-maintenance9-10799)VUID-VkImageMemoryBarrier-maintenance9-10799  
  If the [`maintenance9`](#features-maintenance9) feature is enabled and `image` is a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set and either `subresourceRange.baseArrayLayer` is not equal to 0 or `subresourceRange.layerCount` is not equal to `VK_REMAINING_ARRAY_LAYERS`, `subresourceRange.levelCount` **must** be 1
- [](#VUID-VkImageMemoryBarrier-subresourceRange-01725)VUID-VkImageMemoryBarrier-subresourceRange-01725  
  If `image` is not a 3D image or was created without `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, or the [`maintenance9`](#features-maintenance9) feature is not enabled, and `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `subresourceRange.baseArrayLayer` + `subresourceRange.layerCount` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkImageMemoryBarrier-maintenance9-10800)VUID-VkImageMemoryBarrier-maintenance9-10800  
  If the [`maintenance9`](#features-maintenance9) feature is enabled, `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, and `image` is a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, `subresourceRange.baseArrayLayer` + `subresourceRange.layerCount` **must** be less than or equal to the depth computed from `baseMipLevel` and `extent.depth` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created, according to the formula defined in [Image Mip Level Sizing](#resources-image-mip-level-sizing)
- [](#VUID-VkImageMemoryBarrier-image-01932)VUID-VkImageMemoryBarrier-image-01932  
  If `image` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkImageMemoryBarrier-image-09241)VUID-VkImageMemoryBarrier-image-09241  
  If `image` has a color format that is single-plane, then the `aspectMask` member of `subresourceRange` **must** be `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-VkImageMemoryBarrier-image-09242)VUID-VkImageMemoryBarrier-image-09242  
  If `image` has a color format and is not *disjoint*, then the `aspectMask` member of `subresourceRange` **must** be `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-VkImageMemoryBarrier-image-01672)VUID-VkImageMemoryBarrier-image-01672  
  If `image` has a [multi-planar format](#formats-multiplanar) and the image is *disjoint*, then the `aspectMask` member of `subresourceRange` **must** include at least one [multi-planar aspect mask](#formats-multiplanar-image-aspect) bit or `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-VkImageMemoryBarrier-image-03320)VUID-VkImageMemoryBarrier-image-03320  
  If `image` has a depth/stencil format with both depth and stencil and the [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts) feature is not enabled, then the `aspectMask` member of `subresourceRange` **must** include both `VK_IMAGE_ASPECT_DEPTH_BIT` and `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkImageMemoryBarrier-image-03319)VUID-VkImageMemoryBarrier-image-03319  
  If `image` has a depth/stencil format with both depth and stencil and the [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts) feature is enabled, then the `aspectMask` member of `subresourceRange` **must** include either or both `VK_IMAGE_ASPECT_DEPTH_BIT` and `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkImageMemoryBarrier-image-10749)VUID-VkImageMemoryBarrier-image-10749  
  If `image` has a depth-only format then the `aspectMask` member of `subresourceRange` **must** be `VK_IMAGE_ASPECT_DEPTH_BIT`
- [](#VUID-VkImageMemoryBarrier-image-10750)VUID-VkImageMemoryBarrier-image-10750  
  If `image` has a stencil-only format then the `aspectMask` member of `subresourceRange` **must** be `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkImageMemoryBarrier-aspectMask-08702)VUID-VkImageMemoryBarrier-aspectMask-08702  
  If the `aspectMask` member of `subresourceRange` includes `VK_IMAGE_ASPECT_DEPTH_BIT`, `oldLayout` and `newLayout` **must** not be one of `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkImageMemoryBarrier-aspectMask-08703)VUID-VkImageMemoryBarrier-aspectMask-08703  
  If the `aspectMask` member of `subresourceRange` includes `VK_IMAGE_ASPECT_STENCIL_BIT`, `oldLayout` and `newLayout` **must** not be one of `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`
- [](#VUID-VkImageMemoryBarrier-subresourceRange-09601)VUID-VkImageMemoryBarrier-subresourceRange-09601  
  `subresourceRange.aspectMask` **must** be valid for the `format` the `image` was created with
- [](#VUID-VkImageMemoryBarrier-None-09052)VUID-VkImageMemoryBarrier-None-09052  
  If the [`synchronization2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-synchronization2) feature is not enabled, and `image` was created with a sharing mode of `VK_SHARING_MODE_CONCURRENT`, at least one of `srcQueueFamilyIndex` and `dstQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_IGNORED`
- [](#VUID-VkImageMemoryBarrier-None-09053)VUID-VkImageMemoryBarrier-None-09053  
  If the [`synchronization2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-synchronization2) feature is not enabled, and `image` was created with a sharing mode of `VK_SHARING_MODE_CONCURRENT`, `srcQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_IGNORED` or `VK_QUEUE_FAMILY_EXTERNAL`
- [](#VUID-VkImageMemoryBarrier-None-09054)VUID-VkImageMemoryBarrier-None-09054  
  If the [`synchronization2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-synchronization2) feature is not enabled, and `image` was created with a sharing mode of `VK_SHARING_MODE_CONCURRENT`, `dstQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_IGNORED` or `VK_QUEUE_FAMILY_EXTERNAL`

Valid Usage (Implicit)

- [](#VUID-VkImageMemoryBarrier-sType-sType)VUID-VkImageMemoryBarrier-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER`
- [](#VUID-VkImageMemoryBarrier-pNext-pNext)VUID-VkImageMemoryBarrier-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkExternalMemoryAcquireUnmodifiedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryAcquireUnmodifiedEXT.html) or [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationsInfoEXT.html)
- [](#VUID-VkImageMemoryBarrier-sType-unique)VUID-VkImageMemoryBarrier-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkImageMemoryBarrier-oldLayout-parameter)VUID-VkImageMemoryBarrier-oldLayout-parameter  
  `oldLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkImageMemoryBarrier-newLayout-parameter)VUID-VkImageMemoryBarrier-newLayout-parameter  
  `newLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkImageMemoryBarrier-image-parameter)VUID-VkImageMemoryBarrier-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkImageMemoryBarrier-subresourceRange-parameter)VUID-VkImageMemoryBarrier-subresourceRange-parameter  
  `subresourceRange` **must** be a valid [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAccessFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier.html), [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageMemoryBarrier)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
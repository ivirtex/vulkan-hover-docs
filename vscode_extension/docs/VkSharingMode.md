# VkSharingMode(3) Manual Page

## Name

VkSharingMode - Buffer and image sharing modes



## [](#_c_specification)C Specification

Buffer and image objects are created with a *sharing mode* controlling how they **can** be accessed from queues. The supported sharing modes are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkSharingMode {
    VK_SHARING_MODE_EXCLUSIVE = 0,
    VK_SHARING_MODE_CONCURRENT = 1,
} VkSharingMode;
```

## [](#_description)Description

- `VK_SHARING_MODE_EXCLUSIVE` specifies that access to any range or image subresource of the object will be exclusive to a single queue family at a time.
- `VK_SHARING_MODE_CONCURRENT` specifies that concurrent access to any range or image subresource of the object from multiple queue families is supported.

Note

`VK_SHARING_MODE_CONCURRENT` **may** result in lower performance access to the buffer or image than `VK_SHARING_MODE_EXCLUSIVE`.

Ranges of buffers and image subresources of image objects created using `VK_SHARING_MODE_EXCLUSIVE` **must** only be accessed by queues in the queue family that has *ownership* of the resource. Upon creation, such resources are not owned by any queue family; ownership is implicitly acquired upon first use within a queue. Once a resource using `VK_SHARING_MODE_EXCLUSIVE` is owned by some queue family, unless the [`maintenance9`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance9) feature is enabled, the application **must** perform a [queue family ownership transfer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers) if it wishes to make the memory contents of a range or image subresource accessible to a different queue family. `VK_SHARING_MODE_EXCLUSIVE` resources that are already owned by a queue family **may** be acquired by a different queue family without a queue family ownership transfer, but unless the [`maintenance9`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance9) feature is enabled, their contents become undefined.

If the [`maintenance9`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance9) feature is enabled, the contents of buffer resources, and of linear image resources (i.e., those created with `tiling` set to `VK_IMAGE_TILING_LINEAR`) are always preserved when they are implicitly acquired by a different queue family on the same logical device (i.e., neither queue family is `VK_QUEUE_FAMILY_FOREIGN_EXT` or `VK_QUEUE_FAMILY_EXTERNAL`). This means that whenever the [`maintenance9`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance9) feature is enabled, explicit queue family ownership transfers of buffer and linear image resources between different queue families on the same logical device are **optional**.

Additionally, if the [`maintenance9`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance9) feature is enabled, the contents of some optimal image resources (i.e., those created with `VK_IMAGE_TILING_OPTIMAL`) are always preserved when they are implicitly acquired by a different queue family on the same logical device (i.e., neither queue family is `VK_QUEUE_FAMILY_FOREIGN_EXT` or `VK_QUEUE_FAMILY_EXTERNAL`). This applies only to optimal images that are being implicitly acquired by a queue family whose index bit is set in the current queue family’s [VkQueueFamilyOwnershipTransferPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyOwnershipTransferPropertiesKHR.html)::`optimalImageTransferToQueueFamilies`, and that were created without any of the following bits set in `usage`:

- `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`
- `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- `VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT`
- `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`
- `VK_IMAGE_USAGE_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`
- `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

This means that whenever the [`maintenance9`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance9) feature is enabled, explicit queue family ownership transfers of such image resources between such combinations of queue families are **optional**. For all other optimal images and/or combinations of queue families, the application **must** still perform an explicit queue family ownership transfer if it wishes to make the memory contents of an optimal image subresource already owned by a queue family accessible to a different queue family.

Note

Applications are allowed to perform explicit queue family ownership transfers in circumstances where they are not required, but there is no functional nor performance advantage in doing so. Performing explicit transfers in such cases remains supported for backward compatibility and is not recommended for new applications.

Note

Before being used on the first queue, images still require a [layout transition](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-layouts) from these layouts:

- `VK_IMAGE_LAYOUT_UNDEFINED`
- `VK_IMAGE_LAYOUT_PREINITIALIZED`
- `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT`

A queue family **can** take ownership of an image subresource, tensor subresource, or buffer range of a resource created with `VK_SHARING_MODE_EXCLUSIVE`, without an ownership transfer, in the same way as for a resource that was just created; however, taking ownership in this way has the effect that the contents of the image subresource or buffer range are undefined.

Ranges of buffers, tensor subresources of tensor objects, and image subresources of image objects created using `VK_SHARING_MODE_CONCURRENT` **must** only be accessed by queues from the queue families specified through the `queueFamilyIndexCount` and `pQueueFamilyIndices` members of the corresponding create info structures.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html), [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html), [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html), [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSharingMode).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VkQueueFamilyProperties(3) Manual Page

## Name

VkQueueFamilyProperties - Structure providing information about a queue family



## [](#_c_specification)C Specification

The `VkQueueFamilyProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkQueueFamilyProperties {
    VkQueueFlags    queueFlags;
    uint32_t        queueCount;
    uint32_t        timestampValidBits;
    VkExtent3D      minImageTransferGranularity;
} VkQueueFamilyProperties;
```

## [](#_members)Members

- `queueFlags` is a bitmask of [VkQueueFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFlagBits.html) indicating capabilities of the queues in this queue family.
- `queueCount` is the unsigned integer count of queues in this queue family. Each queue family **must** support at least one queue.
- `timestampValidBits` is the unsigned integer count of meaningful bits in the timestamps written via [vkCmdWriteTimestamp2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteTimestamp2.html) or [vkCmdWriteTimestamp](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteTimestamp.html). The valid range for the count is 36 to 64 bits, or a value of 0, indicating no support for timestamps. Bits outside the valid range are guaranteed to be zeros.
- `minImageTransferGranularity` is the minimum granularity supported for image transfer operations on the queues in this queue family.

## [](#_description)Description

The value returned in `minImageTransferGranularity` has a unit of compressed texel blocks for images having a block-compressed format, and a unit of texels otherwise.

Possible values of `minImageTransferGranularity` are:

- (0,0,0) specifies that only whole mip levels **must** be transferred using the image transfer operations on the corresponding queues. In this case, the following restrictions apply to all offset and extent parameters of image transfer operations:
  
  - The `x`, `y`, and `z` members of a [VkOffset3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset3D.html) parameter **must** always be zero.
  - The `width`, `height`, and `depth` members of a [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html) parameter **must** always match the width, height, and depth of the image subresource corresponding to the parameter, respectively.
- (Ax, Ay, Az) where Ax, Ay, and Az are all integer powers of two. In this case the following restrictions apply to all image transfer operations:
  
  - `x`, `y`, and `z` of a [VkOffset3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset3D.html) parameter **must** be integer multiples of Ax, Ay, and Az, respectively.
  - `width` of a [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html) parameter **must** be an integer multiple of Ax, or else `x` + `width` **must** equal the width of the image subresource corresponding to the parameter.
  - `height` of a [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html) parameter **must** be an integer multiple of Ay, or else `y` + `height` **must** equal the height of the image subresource corresponding to the parameter.
  - `depth` of a [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html) parameter **must** be an integer multiple of Az, or else `z` + `depth` **must** equal the depth of the image subresource corresponding to the parameter.
  - If the format of the image corresponding to the parameters is one of the block-compressed formats then for the purposes of the above calculations the granularity **must** be scaled up by the compressed texel block dimensions.

Queues supporting graphics and/or compute operations **must** report (1,1,1) in `minImageTransferGranularity`, meaning that there are no additional restrictions on the granularity of image transfer operations for these queues. Other queues supporting image transfer operations are only **required** to support whole mip level transfers, thus `minImageTransferGranularity` for queues belonging to such queue families **may** be (0,0,0).

The [Device Memory](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-device) section describes memory properties queried from the physical device.

For physical device feature queries see the [Features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features) chapter.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html), [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties2.html), [VkQueueFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFlags.html), [vkGetPhysicalDeviceQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueueFamilyProperties).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
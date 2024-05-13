# VkQueueFamilyProperties(3) Manual Page

## Name

VkQueueFamilyProperties - Structure providing information about a queue
family



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkQueueFamilyProperties` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkQueueFamilyProperties {
    VkQueueFlags    queueFlags;
    uint32_t        queueCount;
    uint32_t        timestampValidBits;
    VkExtent3D      minImageTransferGranularity;
} VkQueueFamilyProperties;
```

## <a href="#_members" class="anchor"></a>Members

- `queueFlags` is a bitmask of [VkQueueFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFlagBits.html)
  indicating capabilities of the queues in this queue family.

- `queueCount` is the unsigned integer count of queues in this queue
  family. Each queue family **must** support at least one queue.

- `timestampValidBits` is the unsigned integer count of meaningful bits
  in the timestamps written via
  [vkCmdWriteTimestamp2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteTimestamp2.html) or
  [vkCmdWriteTimestamp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteTimestamp.html). The valid range for
  the count is 36 to 64 bits, or a value of 0, indicating no support for
  timestamps. Bits outside the valid range are guaranteed to be zeros.

- `minImageTransferGranularity` is the minimum granularity supported for
  image transfer operations on the queues in this queue family.

## <a href="#_description" class="anchor"></a>Description

The value returned in `minImageTransferGranularity` has a unit of
compressed texel blocks for images having a block-compressed format, and
a unit of texels otherwise.

Possible values of `minImageTransferGranularity` are:

- (0,0,0) specifies that only whole mip levels **must** be transferred
  using the image transfer operations on the corresponding queues. In
  this case, the following restrictions apply to all offset and extent
  parameters of image transfer operations:

  - The `x`, `y`, and `z` members of a [VkOffset3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset3D.html)
    parameter **must** always be zero.

  - The `width`, `height`, and `depth` members of a
    [VkExtent3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html) parameter **must** always match the
    width, height, and depth of the image subresource corresponding to
    the parameter, respectively.

- (A<sub>x</sub>, A<sub>y</sub>, A<sub>z</sub>) where A<sub>x</sub>,
  A<sub>y</sub>, and A<sub>z</sub> are all integer powers of two. In
  this case the following restrictions apply to all image transfer
  operations:

  - `x`, `y`, and `z` of a [VkOffset3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset3D.html) parameter
    **must** be integer multiples of A<sub>x</sub>, A<sub>y</sub>, and
    A<sub>z</sub>, respectively.

  - `width` of a [VkExtent3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html) parameter **must** be an
    integer multiple of A<sub>x</sub>, or else `x` + `width` **must**
    equal the width of the image subresource corresponding to the
    parameter.

  - `height` of a [VkExtent3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html) parameter **must** be an
    integer multiple of A<sub>y</sub>, or else `y` + `height` **must**
    equal the height of the image subresource corresponding to the
    parameter.

  - `depth` of a [VkExtent3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html) parameter **must** be an
    integer multiple of A<sub>z</sub>, or else `z` + `depth` **must**
    equal the depth of the image subresource corresponding to the
    parameter.

  - If the format of the image corresponding to the parameters is one of
    the block-compressed formats then for the purposes of the above
    calculations the granularity **must** be scaled up by the compressed
    texel block dimensions.

Queues supporting graphics and/or compute operations **must** report
(1,1,1) in `minImageTransferGranularity`, meaning that there are no
additional restrictions on the granularity of image transfer operations
for these queues. Other queues supporting image transfer operations are
only **required** to support whole mip level transfers, thus
`minImageTransferGranularity` for queues belonging to such queue
families **may** be (0,0,0).

The <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-device"
target="_blank" rel="noopener">Device Memory</a> section describes
memory properties queried from the physical device.

For physical device feature queries see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features"
target="_blank" rel="noopener">Features</a> chapter.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html),
[VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2.html),
[VkQueueFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFlags.html),
[vkGetPhysicalDeviceQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueueFamilyProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

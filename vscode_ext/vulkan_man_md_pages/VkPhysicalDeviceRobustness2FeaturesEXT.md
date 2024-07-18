# VkPhysicalDeviceRobustness2FeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceRobustness2FeaturesEXT - Structure describing the
out-of-bounds behavior for an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceRobustness2FeaturesEXT` structure is defined as:

``` c
// Provided by VK_EXT_robustness2
typedef struct VkPhysicalDeviceRobustness2FeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           robustBufferAccess2;
    VkBool32           robustImageAccess2;
    VkBool32           nullDescriptor;
} VkPhysicalDeviceRobustness2FeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-robustBufferAccess2"></span> `robustBufferAccess2`
  indicates whether buffer accesses are tightly bounds-checked against
  the range of the descriptor. Uniform buffers **must** be
  bounds-checked to the range of the descriptor, where the range is
  rounded up to a multiple of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-robustUniformBufferAccessSizeAlignment"
  target="_blank"
  rel="noopener"><code>robustUniformBufferAccessSizeAlignment</code></a>.
  Storage buffers **must** be bounds-checked to the range of the
  descriptor, where the range is rounded up to a multiple of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-robustStorageBufferAccessSizeAlignment"
  target="_blank"
  rel="noopener"><code>robustStorageBufferAccessSizeAlignment</code></a>.
  Out of bounds buffer loads will return zero values, and [image load,
  sample, and atomic operations](#textures) from texel buffers will have
  (0,0,1) values [inserted for missing G, B, or A
  components](#textures-conversion-to-rgba) based on the format.

- <span id="features-robustImageAccess2"></span> `robustImageAccess2`
  indicates whether image accesses are tightly bounds-checked against
  the dimensions of the image view. Out of bounds [image load, sample,
  and atomic operations](#textures) from images will return zero values,
  with (0,0,1) values [inserted for missing G, B, or A
  components](#textures-conversion-to-rgba) based on the format.

- <span id="features-nullDescriptor"></span> `nullDescriptor` indicates
  whether descriptors **can** be written with a
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) resource or view, which are
  considered valid to access and act as if the descriptor were bound to
  nothing.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceRobustness2FeaturesEXT` structure is included in
the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceRobustness2FeaturesEXT` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage

- <a
  href="#VUID-VkPhysicalDeviceRobustness2FeaturesEXT-robustBufferAccess2-04000"
  id="VUID-VkPhysicalDeviceRobustness2FeaturesEXT-robustBufferAccess2-04000"></a>
  VUID-VkPhysicalDeviceRobustness2FeaturesEXT-robustBufferAccess2-04000  
  If `robustBufferAccess2` is enabled then <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess"
  target="_blank" rel="noopener"><code>robustBufferAccess</code></a>
  **must** also be enabled

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceRobustness2FeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceRobustness2FeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceRobustness2FeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ROBUSTNESS_2_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_robustness2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_robustness2.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceRobustness2FeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

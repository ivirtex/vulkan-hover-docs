# VkPhysicalDeviceUniformBufferStandardLayoutFeatures(3) Manual Page

## Name

VkPhysicalDeviceUniformBufferStandardLayoutFeatures - Structure
indicating support for std430-like packing in uniform buffers



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceUniformBufferStandardLayoutFeatures` structure is
defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceUniformBufferStandardLayoutFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           uniformBufferStandardLayout;
} VkPhysicalDeviceUniformBufferStandardLayoutFeatures;
```

or the equivalent

``` c
// Provided by VK_KHR_uniform_buffer_standard_layout
typedef VkPhysicalDeviceUniformBufferStandardLayoutFeatures VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-uniformBufferStandardLayout"></span>
  `uniformBufferStandardLayout` indicates that the implementation
  supports the same layouts for uniform buffers as for storage and other
  kinds of buffers. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-resources-standard-layout"
  target="_blank" rel="noopener">Standard Buffer Layout</a>.

If the `VkPhysicalDeviceUniformBufferStandardLayoutFeatures` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceUniformBufferStandardLayoutFeatures` **can** also be
used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceUniformBufferStandardLayoutFeatures-sType-sType"
  id="VUID-VkPhysicalDeviceUniformBufferStandardLayoutFeatures-sType-sType"></a>
  VUID-VkPhysicalDeviceUniformBufferStandardLayoutFeatures-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_uniform_buffer_standard_layout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_uniform_buffer_standard_layout.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceUniformBufferStandardLayoutFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

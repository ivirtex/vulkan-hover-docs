# VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR - Structure
indicating support for local reads in dynamic render pass instances



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR` structure is
defined as:

``` c
// Provided by VK_KHR_dynamic_rendering_local_read
typedef struct VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           dynamicRenderingLocalRead;
} VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-dynamicRenderingLocalRead"></span>
  `dynamicRenderingLocalRead` specifies that the implementation supports
  local reads inside dynamic render pass instances using the
  [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html) command.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR` **can** also be
used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_LOCAL_READ_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_dynamic_rendering_local_read](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering_local_read.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

# VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM(3) Manual Page

## Name

VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM - Structure
describing multiview per view render areas features that can be
supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM` structure
is defined as:

``` c
// Provided by VK_QCOM_multiview_per_view_render_areas
typedef struct VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           multiviewPerViewRenderAreas;
} VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- <span id="features-multiview-per-view-render-areas"></span>
  `multiviewPerViewRenderAreas` indicates that the implementation
  supports multiview per-view render areas.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM` **can** also
be used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM-sType-sType"
  id="VUID-VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM-sType-sType"></a>
  VUID-VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PER_VIEW_RENDER_AREAS_FEATURES_QCOM`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_multiview_per_view_render_areas](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_multiview_per_view_render_areas.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

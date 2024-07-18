# vkGetRenderingAreaGranularityKHR(3) Manual Page

## Name

vkGetRenderingAreaGranularityKHR - Returns the granularity for dynamic
rendering optimal render area



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the render area granularity for a render pass instance, call:

``` c
// Provided by VK_KHR_maintenance5
void vkGetRenderingAreaGranularityKHR(
    VkDevice                                    device,
    const VkRenderingAreaInfoKHR*               pRenderingAreaInfo,
    VkExtent2D*                                 pGranularity);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the render pass instance.

- `pRenderingAreaInfo` is a pointer to a
  [VkRenderingAreaInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAreaInfoKHR.html) structure
  specifying details of the render pass instance to query the render
  area granularity for.

- `pGranularity` is a pointer to a [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html)
  structure in which the granularity is returned.

## <a href="#_description" class="anchor"></a>Description

The conditions leading to an optimal `renderArea` are:

- the `offset.x` member in `renderArea` is a multiple of the `width`
  member of the returned [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html) (the horizontal
  granularity).

- the `offset.y` member in `renderArea` is a multiple of the `height`
  member of the returned [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html) (the vertical
  granularity).

- either the `extent.width` member in `renderArea` is a multiple of the
  horizontal granularity or `offset.x`+`extent.width` is equal to the
  `width` of each attachment used in the render pass instance.

- either the `extent.height` member in `renderArea` is a multiple of the
  vertical granularity or `offset.y`+`extent.height` is equal to the
  `height` of each attachment used in the render pass instance.

Valid Usage (Implicit)

- <a href="#VUID-vkGetRenderingAreaGranularityKHR-device-parameter"
  id="VUID-vkGetRenderingAreaGranularityKHR-device-parameter"></a>
  VUID-vkGetRenderingAreaGranularityKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetRenderingAreaGranularityKHR-pRenderingAreaInfo-parameter"
  id="VUID-vkGetRenderingAreaGranularityKHR-pRenderingAreaInfo-parameter"></a>
  VUID-vkGetRenderingAreaGranularityKHR-pRenderingAreaInfo-parameter  
  `pRenderingAreaInfo` **must** be a valid pointer to a valid
  [VkRenderingAreaInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAreaInfoKHR.html) structure

- <a href="#VUID-vkGetRenderingAreaGranularityKHR-pGranularity-parameter"
  id="VUID-vkGetRenderingAreaGranularityKHR-pGranularity-parameter"></a>
  VUID-vkGetRenderingAreaGranularityKHR-pGranularity-parameter  
  `pGranularity` **must** be a valid pointer to a
  [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance5](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance5.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html),
[VkRenderingAreaInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAreaInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetRenderingAreaGranularityKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

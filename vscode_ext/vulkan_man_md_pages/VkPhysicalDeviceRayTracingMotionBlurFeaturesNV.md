# VkPhysicalDeviceRayTracingMotionBlurFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceRayTracingMotionBlurFeaturesNV - Structure describing
the ray tracing motion blur features that can be supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceRayTracingMotionBlurFeaturesNV` structure is
defined as:

``` c
// Provided by VK_NV_ray_tracing_motion_blur
typedef struct VkPhysicalDeviceRayTracingMotionBlurFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           rayTracingMotionBlur;
    VkBool32           rayTracingMotionBlurPipelineTraceRaysIndirect;
} VkPhysicalDeviceRayTracingMotionBlurFeaturesNV;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-rayTracingMotionBlur"></span>
  `rayTracingMotionBlur` indicates whether the implementation supports
  the motion blur feature.

- <span id="features-rayTracingMotionBlurPipelineTraceRaysIndirect"></span>
  `rayTracingMotionBlurPipelineTraceRaysIndirect` indicates whether the
  implementation supports indirect ray tracing commands with the motion
  blur feature enabled.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceRayTracingMotionBlurFeaturesNV` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceRayTracingMotionBlurFeaturesNV` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceRayTracingMotionBlurFeaturesNV-sType-sType"
  id="VUID-VkPhysicalDeviceRayTracingMotionBlurFeaturesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceRayTracingMotionBlurFeaturesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_TRACING_MOTION_BLUR_FEATURES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing_motion_blur](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing_motion_blur.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceRayTracingMotionBlurFeaturesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

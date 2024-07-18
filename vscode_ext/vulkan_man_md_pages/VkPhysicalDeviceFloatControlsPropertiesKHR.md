# VkPhysicalDeviceFloatControlsProperties(3) Manual Page

## Name

VkPhysicalDeviceFloatControlsProperties - Structure describing
properties supported by VK_KHR_shader_float_controls



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceFloatControlsProperties` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceFloatControlsProperties {
    VkStructureType                      sType;
    void*                                pNext;
    VkShaderFloatControlsIndependence    denormBehaviorIndependence;
    VkShaderFloatControlsIndependence    roundingModeIndependence;
    VkBool32                             shaderSignedZeroInfNanPreserveFloat16;
    VkBool32                             shaderSignedZeroInfNanPreserveFloat32;
    VkBool32                             shaderSignedZeroInfNanPreserveFloat64;
    VkBool32                             shaderDenormPreserveFloat16;
    VkBool32                             shaderDenormPreserveFloat32;
    VkBool32                             shaderDenormPreserveFloat64;
    VkBool32                             shaderDenormFlushToZeroFloat16;
    VkBool32                             shaderDenormFlushToZeroFloat32;
    VkBool32                             shaderDenormFlushToZeroFloat64;
    VkBool32                             shaderRoundingModeRTEFloat16;
    VkBool32                             shaderRoundingModeRTEFloat32;
    VkBool32                             shaderRoundingModeRTEFloat64;
    VkBool32                             shaderRoundingModeRTZFloat16;
    VkBool32                             shaderRoundingModeRTZFloat32;
    VkBool32                             shaderRoundingModeRTZFloat64;
} VkPhysicalDeviceFloatControlsProperties;
```

or the equivalent

``` c
// Provided by VK_KHR_shader_float_controls
typedef VkPhysicalDeviceFloatControlsProperties VkPhysicalDeviceFloatControlsPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-denormBehaviorIndependence"></span>
  `denormBehaviorIndependence` is a
  [VkShaderFloatControlsIndependence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderFloatControlsIndependence.html)
  value indicating whether, and how, denorm behavior can be set
  independently for different bit widths.

- <span id="extension-features-roundingModeIndependence"></span>
  `roundingModeIndependence` is a
  [VkShaderFloatControlsIndependence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderFloatControlsIndependence.html)
  value indicating whether, and how, rounding modes can be set
  independently for different bit widths.

- <span id="extension-limits-shaderSignedZeroInfNanPreserveFloat16"></span>
  `shaderSignedZeroInfNanPreserveFloat16` is a boolean value indicating
  whether sign of a zero, Nans and ±∞ **can** be preserved in 16-bit
  floating-point computations. It also indicates whether the
  `SignedZeroInfNanPreserve` execution mode **can** be used for 16-bit
  floating-point types.

- <span id="extension-limits-shaderSignedZeroInfNanPreserveFloat32"></span>
  `shaderSignedZeroInfNanPreserveFloat32` is a boolean value indicating
  whether sign of a zero, Nans and ±∞ **can** be preserved in 32-bit
  floating-point computations. It also indicates whether the
  `SignedZeroInfNanPreserve` execution mode **can** be used for 32-bit
  floating-point types.

- <span id="extension-limits-shaderSignedZeroInfNanPreserveFloat64"></span>
  `shaderSignedZeroInfNanPreserveFloat64` is a boolean value indicating
  whether sign of a zero, Nans and ±∞ **can** be preserved in 64-bit
  floating-point computations. It also indicates whether the
  `SignedZeroInfNanPreserve` execution mode **can** be used for 64-bit
  floating-point types.

- <span id="extension-limits-shaderDenormPreserveFloat16"></span>
  `shaderDenormPreserveFloat16` is a boolean value indicating whether
  denormals **can** be preserved in 16-bit floating-point computations.
  It also indicates whether the `DenormPreserve` execution mode **can**
  be used for 16-bit floating-point types.

- <span id="extension-limits-shaderDenormPreserveFloat32"></span>
  `shaderDenormPreserveFloat32` is a boolean value indicating whether
  denormals **can** be preserved in 32-bit floating-point computations.
  It also indicates whether the `DenormPreserve` execution mode **can**
  be used for 32-bit floating-point types.

- <span id="extension-limits-shaderDenormPreserveFloat64"></span>
  `shaderDenormPreserveFloat64` is a boolean value indicating whether
  denormals **can** be preserved in 64-bit floating-point computations.
  It also indicates whether the `DenormPreserve` execution mode **can**
  be used for 64-bit floating-point types.

- <span id="extension-limits-shaderDenormFlushToZeroFloat16"></span>
  `shaderDenormFlushToZeroFloat16` is a boolean value indicating whether
  denormals **can** be flushed to zero in 16-bit floating-point
  computations. It also indicates whether the `DenormFlushToZero`
  execution mode **can** be used for 16-bit floating-point types.

- <span id="extension-limits-shaderDenormFlushToZeroFloat32"></span>
  `shaderDenormFlushToZeroFloat32` is a boolean value indicating whether
  denormals **can** be flushed to zero in 32-bit floating-point
  computations. It also indicates whether the `DenormFlushToZero`
  execution mode **can** be used for 32-bit floating-point types.

- <span id="extension-limits-shaderDenormFlushToZeroFloat64"></span>
  `shaderDenormFlushToZeroFloat64` is a boolean value indicating whether
  denormals **can** be flushed to zero in 64-bit floating-point
  computations. It also indicates whether the `DenormFlushToZero`
  execution mode **can** be used for 64-bit floating-point types.

- <span id="extension-limits-shaderRoundingModeRTEFloat16"></span>
  `shaderRoundingModeRTEFloat16` is a boolean value indicating whether
  an implementation supports the round-to-nearest-even rounding mode for
  16-bit floating-point arithmetic and conversion instructions. It also
  indicates whether the `RoundingModeRTE` execution mode **can** be used
  for 16-bit floating-point types.

- <span id="extension-limits-shaderRoundingModeRTEFloat32"></span>
  `shaderRoundingModeRTEFloat32` is a boolean value indicating whether
  an implementation supports the round-to-nearest-even rounding mode for
  32-bit floating-point arithmetic and conversion instructions. It also
  indicates whether the `RoundingModeRTE` execution mode **can** be used
  for 32-bit floating-point types.

- <span id="extension-limits-shaderRoundingModeRTEFloat64"></span>
  `shaderRoundingModeRTEFloat64` is a boolean value indicating whether
  an implementation supports the round-to-nearest-even rounding mode for
  64-bit floating-point arithmetic and conversion instructions. It also
  indicates whether the `RoundingModeRTE` execution mode **can** be used
  for 64-bit floating-point types.

- <span id="extension-limits-shaderRoundingModeRTZFloat16"></span>
  `shaderRoundingModeRTZFloat16` is a boolean value indicating whether
  an implementation supports the round-towards-zero rounding mode for
  16-bit floating-point arithmetic and conversion instructions. It also
  indicates whether the `RoundingModeRTZ` execution mode **can** be used
  for 16-bit floating-point types.

- <span id="extension-limits-shaderRoundingModeRTZFloat32"></span>
  `shaderRoundingModeRTZFloat32` is a boolean value indicating whether
  an implementation supports the round-towards-zero rounding mode for
  32-bit floating-point arithmetic and conversion instructions. It also
  indicates whether the `RoundingModeRTZ` execution mode **can** be used
  for 32-bit floating-point types.

- <span id="extension-limits-shaderRoundingModeRTZFloat64"></span>
  `shaderRoundingModeRTZFloat64` is a boolean value indicating whether
  an implementation supports the round-towards-zero rounding mode for
  64-bit floating-point arithmetic and conversion instructions. It also
  indicates whether the `RoundingModeRTZ` execution mode **can** be used
  for 64-bit floating-point types.

If the `VkPhysicalDeviceFloatControlsProperties` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceFloatControlsProperties-sType-sType"
  id="VUID-VkPhysicalDeviceFloatControlsProperties-sType-sType"></a>
  VUID-VkPhysicalDeviceFloatControlsProperties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_shader_float_controls](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_float_controls.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkShaderFloatControlsIndependence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderFloatControlsIndependence.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceFloatControlsProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

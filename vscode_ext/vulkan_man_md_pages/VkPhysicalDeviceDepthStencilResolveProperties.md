# VkPhysicalDeviceDepthStencilResolveProperties(3) Manual Page

## Name

VkPhysicalDeviceDepthStencilResolveProperties - Structure describing
depth/stencil resolve properties that can be supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceDepthStencilResolveProperties` structure is defined
as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceDepthStencilResolveProperties {
    VkStructureType       sType;
    void*                 pNext;
    VkResolveModeFlags    supportedDepthResolveModes;
    VkResolveModeFlags    supportedStencilResolveModes;
    VkBool32              independentResolveNone;
    VkBool32              independentResolve;
} VkPhysicalDeviceDepthStencilResolveProperties;
```

or the equivalent

``` c
// Provided by VK_KHR_depth_stencil_resolve
typedef VkPhysicalDeviceDepthStencilResolveProperties VkPhysicalDeviceDepthStencilResolvePropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-depthResolveModes"></span>
  `supportedDepthResolveModes` is a bitmask of
  [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html) indicating the set
  of supported depth resolve modes. `VK_RESOLVE_MODE_SAMPLE_ZERO_BIT`
  **must** be included in the set but implementations **may** support
  additional modes.

- <span id="extension-features-stencilResolveModes"></span>
  `supportedStencilResolveModes` is a bitmask of
  [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html) indicating the set
  of supported stencil resolve modes. `VK_RESOLVE_MODE_SAMPLE_ZERO_BIT`
  **must** be included in the set but implementations **may** support
  additional modes. `VK_RESOLVE_MODE_AVERAGE_BIT` **must** not be
  included in the set.

- <span id="extension-features-independentResolveNone"></span>
  `independentResolveNone` is `VK_TRUE` if the implementation supports
  setting the depth and stencil resolve modes to different values when
  one of those modes is `VK_RESOLVE_MODE_NONE`. Otherwise the
  implementation only supports setting both modes to the same value.

- <span id="extension-features-independentResolve"></span>
  `independentResolve` is `VK_TRUE` if the implementation supports all
  combinations of the supported depth and stencil resolve modes,
  including setting either depth or stencil resolve mode to
  `VK_RESOLVE_MODE_NONE`. An implementation that supports
  `independentResolve` **must** also support `independentResolveNone`.

If the `VkPhysicalDeviceDepthStencilResolveProperties` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceDepthStencilResolveProperties-sType-sType"
  id="VUID-VkPhysicalDeviceDepthStencilResolveProperties-sType-sType"></a>
  VUID-VkPhysicalDeviceDepthStencilResolveProperties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_depth_stencil_resolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_depth_stencil_resolve.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkResolveModeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceDepthStencilResolveProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

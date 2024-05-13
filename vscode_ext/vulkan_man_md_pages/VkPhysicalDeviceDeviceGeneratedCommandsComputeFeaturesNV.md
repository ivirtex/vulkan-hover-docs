# VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV(3) Manual Page

## Name

VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV - Structure
describing the device-generated compute features that can be supported
by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV` structure
is defined as:

``` c
// Provided by VK_NV_device_generated_commands_compute
typedef struct VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           deviceGeneratedCompute;
    VkBool32           deviceGeneratedComputePipelines;
    VkBool32           deviceGeneratedComputeCaptureReplay;
} VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-deviceGeneratedCompute"></span>
  `deviceGeneratedCompute` indicates whether the implementation supports
  functionality to generate dispatch commands and push constants for the
  compute pipeline on the device. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#device-generated-commands"
  target="_blank" rel="noopener">Device-Generated Commands</a>.

- <span id="features-deviceGeneratedComputePipelines"></span>
  `deviceGeneratedComputePipelines` indicates whether the implementation
  supports functionality to generate commands to bind compute pipelines
  on the device. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#device-generated-commands"
  target="_blank" rel="noopener">Device-Generated Commands</a>.

- <span id="features-deviceGeneratedComputeCaptureReplay"></span>
  `deviceGeneratedComputeCaptureReplay` indicates whether the
  implementation supports functionality to capture compute pipeline
  address and reuse later for replay in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#device-generated-commands"
  target="_blank" rel="noopener">Device-Generated Commands</a>.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV` **can** also
be used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV-sType-sType"
  id="VUID-VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEVICE_GENERATED_COMMANDS_COMPUTE_FEATURES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands_compute](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands_compute.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

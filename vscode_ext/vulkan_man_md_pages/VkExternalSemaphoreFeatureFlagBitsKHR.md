# VkExternalSemaphoreFeatureFlagBits(3) Manual Page

## Name

VkExternalSemaphoreFeatureFlagBits - Bitfield describing features of an
external semaphore handle type



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in
[VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreProperties.html)::`externalSemaphoreFeatures`,
specifying the features of an external semaphore handle type, are:

``` c
// Provided by VK_VERSION_1_1
typedef enum VkExternalSemaphoreFeatureFlagBits {
    VK_EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT = 0x00000001,
    VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT = 0x00000002,
  // Provided by VK_KHR_external_semaphore_capabilities
    VK_EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT_KHR = VK_EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT,
  // Provided by VK_KHR_external_semaphore_capabilities
    VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT_KHR = VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT,
} VkExternalSemaphoreFeatureFlagBits;
```

or the equivalent

``` c
// Provided by VK_KHR_external_semaphore_capabilities
typedef VkExternalSemaphoreFeatureFlagBits VkExternalSemaphoreFeatureFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT` specifies that handles
  of this type **can** be exported from Vulkan semaphore objects.

- `VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT` specifies that handles
  of this type **can** be imported as Vulkan semaphore objects.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalSemaphoreFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreFeatureFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExternalSemaphoreFeatureFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

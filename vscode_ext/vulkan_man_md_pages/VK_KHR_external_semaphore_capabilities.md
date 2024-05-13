# VK_KHR_external_semaphore_capabilities(3) Manual Page

## Name

VK_KHR_external_semaphore_capabilities - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

77

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_semaphore_capabilities%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_semaphore_capabilities%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-10-20

**IP Status**  
No known IP claims.

**Contributors**  
- Jesse Hall, Google

- James Jones, NVIDIA

- Jeff Juliano, NVIDIA

## <a href="#_description" class="anchor"></a>Description

An application may wish to reference device semaphores in multiple
Vulkan logical devices or instances, in multiple processes, and/or in
multiple APIs. This extension provides a set of capability queries and
handle definitions that allow an application to determine what types of
“external” semaphore handles an implementation supports for a given set
of use cases.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetPhysicalDeviceExternalSemaphorePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalSemaphorePropertiesKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkExternalSemaphorePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphorePropertiesKHR.html)

- [VkPhysicalDeviceExternalSemaphoreInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalSemaphoreInfoKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceIDPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceIDPropertiesKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkExternalSemaphoreFeatureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreFeatureFlagBitsKHR.html)

- [VkExternalSemaphoreHandleTypeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBitsKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkExternalSemaphoreFeatureFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreFeatureFlagsKHR.html)

- [VkExternalSemaphoreHandleTypeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_EXTERNAL_SEMAPHORE_CAPABILITIES_EXTENSION_NAME`

- `VK_KHR_EXTERNAL_SEMAPHORE_CAPABILITIES_SPEC_VERSION`

- `VK_LUID_SIZE_KHR`

- Extending
  [VkExternalSemaphoreFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreFeatureFlagBits.html):

  - `VK_EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT_KHR`

  - `VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT_KHR`

- Extending
  [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html):

  - `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT_KHR`

  - `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR`

  - `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR`

  - `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR`

  - `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-10-20 (James Jones)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_external_semaphore_capabilities"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

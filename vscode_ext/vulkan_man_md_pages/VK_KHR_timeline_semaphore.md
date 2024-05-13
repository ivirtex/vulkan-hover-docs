# VK_KHR_timeline_semaphore(3) Manual Page

## Name

VK_KHR_timeline_semaphore - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

208

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Faith Ekstrand <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_timeline_semaphore%5D%20@gfxstrand%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_timeline_semaphore%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>gfxstrand</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-06-12

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension interacts with
  [`VK_KHR_external_semaphore_capabilities`](VK_KHR_external_semaphore_capabilities.html)

- This extension interacts with
  [`VK_KHR_external_semaphore`](VK_KHR_external_semaphore.html)

- This extension interacts with
  [`VK_KHR_external_semaphore_win32`](VK_KHR_external_semaphore_win32.html)

**Contributors**  
- Jeff Bolz, NVIDIA

- Yuriy O’Donnell, Epic Games

- Faith Ekstrand, Intel

- Jesse Hall, Google

- James Jones, NVIDIA

- Jeff Juliano, NVIDIA

- Daniel Rakos, AMD

- Ray Smith, Arm

## <a href="#_description" class="anchor"></a>Description

This extension introduces a new type of semaphore that has an integer
payload identifying a point in a timeline. Such timeline semaphores
support the following operations:

- Host query - A host operation that allows querying the payload of the
  timeline semaphore.

- Host wait - A host operation that allows a blocking wait for a
  timeline semaphore to reach a specified value.

- Host signal - A host operation that allows advancing the timeline
  semaphore to a specified value.

- Device wait - A device operation that allows waiting for a timeline
  semaphore to reach a specified value.

- Device signal - A device operation that allows advancing the timeline
  semaphore to a specified value.

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetSemaphoreCounterValueKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreCounterValueKHR.html)

- [vkSignalSemaphoreKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSignalSemaphoreKHR.html)

- [vkWaitSemaphoresKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWaitSemaphoresKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkSemaphoreSignalInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSignalInfoKHR.html)

- [VkSemaphoreWaitInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitInfoKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceTimelineSemaphoreFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTimelineSemaphoreFeaturesKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceTimelineSemaphorePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTimelineSemaphorePropertiesKHR.html)

- Extending [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html),
  [VkPhysicalDeviceExternalSemaphoreInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalSemaphoreInfo.html):

  - [VkSemaphoreTypeCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfoKHR.html)

- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html),
  [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindSparseInfo.html):

  - [VkTimelineSemaphoreSubmitInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfoKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkSemaphoreTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeKHR.html)

- [VkSemaphoreWaitFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitFlagBitsKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkSemaphoreWaitFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_TIMELINE_SEMAPHORE_EXTENSION_NAME`

- `VK_KHR_TIMELINE_SEMAPHORE_SPEC_VERSION`

- Extending [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html):

  - `VK_SEMAPHORE_TYPE_BINARY_KHR`

  - `VK_SEMAPHORE_TYPE_TIMELINE_KHR`

- Extending [VkSemaphoreWaitFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitFlagBits.html):

  - `VK_SEMAPHORE_WAIT_ANY_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_SEMAPHORE_SIGNAL_INFO_KHR`

  - `VK_STRUCTURE_TYPE_SEMAPHORE_TYPE_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_SEMAPHORE_WAIT_INFO_KHR`

  - `VK_STRUCTURE_TYPE_TIMELINE_SEMAPHORE_SUBMIT_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Do we need a new object type for this?

**RESOLVED**: No, we just introduce a new type of semaphore object, as
`VK_KHR_external_semaphore_win32` already uses semaphores as the
destination for importing D3D12 fence objects, which are semantically
close/identical to the proposed synchronization primitive.

2\) What type of payload the new synchronization primitive has?

**RESOLVED**: A 64-bit unsigned integer that can only be set to strictly
increasing values by signal operations and is not changed by wait
operations.

3\) Does the new synchronization primitive have the same
signal-before-wait requirement as the existing semaphores do?

**RESOLVED**: No. Timeline semaphores support signaling and waiting
entirely asynchronously. It is the responsibility of the client to avoid
deadlock.

4\) Does the new synchronization primitive allow resetting its payload?

**RESOLVED**: No, allowing the payload value to “go backwards” is
problematic. Applications looking for reset behavior should create a new
instance of the synchronization primitive instead.

5\) How do we enable host waits on the synchronization primitive?

**RESOLVED**: Both a non-blocking query of the current payload value of
the synchronization primitive, and a blocking wait operation are
provided.

6\) How do we enable device waits and signals on the synchronization
primitive?

**RESOLVED**: Similar to `VK_KHR_external_semaphore_win32`, this
extension introduces a new structure that can be chained to
[VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html) to specify the values signaled
semaphores should be set to, and the values waited semaphores need to
reach.

7\) Can the new synchronization primitive be used to synchronize
presentation and swapchain image acquisition operations?

**RESOLVED**: Some implementations may have problems with supporting
that directly, thus it is not allowed in this extension.

8\) Do we want to support external sharing of the new synchronization
primitive type?

**RESOLVED**: Yes. Timeline semaphore specific external sharing
capabilities can be queried using
[vkGetPhysicalDeviceExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalSemaphoreProperties.html)
by chaining the new
[VkSemaphoreTypeCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfoKHR.html)
structure to its `pExternalSemaphoreInfo` structure. This allows having
a different set of external semaphore handle types supported for
timeline semaphores vs. binary semaphores.

9\) Do we need to add a host signal operation for the new
synchronization primitive type?

**RESOLVED**: Yes. This helps in situations where one host thread
submits a workload but another host thread has the information on when
the workload is ready to be executed.

10\) How should the new synchronization primitive interact with the
ordering requirements of the original `VkSemaphore`?

**RESOLVED**: Prior to calling any command which **may** cause a wait
operation on a binary semaphore, the client **must** ensure that the
semaphore signal operation that has been submitted for execution and any
semaphore signal operations on which it depends (if any) **must** have
also been submitted for execution.

11\) Should we have separate feature bits for different sub-features of
timeline semaphores?

**RESOLVED**: No. The only feature which cannot be supported universally
is timeline semaphore import/export. For import/export, the client is
already required to query available external handle types via
[vkGetPhysicalDeviceExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalSemaphoreProperties.html)
and provide the semaphore type by adding a
[VkSemaphoreTypeCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfoKHR.html)
structure to the `pNext` chain of
[VkPhysicalDeviceExternalSemaphoreInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalSemaphoreInfo.html)
so no new feature bit is required.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-05-10 (Faith Ekstrand)

  - Initial version

- Revision 2, 2019-06-12 (Faith Ekstrand)

  - Added an initialValue parameter to timeline semaphore creation

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_timeline_semaphore"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

# VK_KHR_external_memory(3) Manual Page

## Name

VK_KHR_external_memory - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

73

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_external_memory_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_capabilities.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_memory%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_memory%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-10-20

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- Interacts with
  [`VK_KHR_dedicated_allocation`](VK_KHR_dedicated_allocation.html).

- Interacts with
  [`VK_NV_dedicated_allocation`](VK_NV_dedicated_allocation.html).

**Contributors**  
- Faith Ekstrand, Intel

- Ian Elliott, Google

- Jesse Hall, Google

- Tobias Hector, Imagination Technologies

- James Jones, NVIDIA

- Jeff Juliano, NVIDIA

- Matthew Netsch, Qualcomm Technologies, Inc.

- Daniel Rakos, AMD

- Carsten Rohde, NVIDIA

- Ray Smith, ARM

- Lina Versace, Google

## <a href="#_description" class="anchor"></a>Description

An application may wish to reference device memory in multiple Vulkan
logical devices or instances, in multiple processes, and/or in multiple
APIs. This extension enables an application to export non-Vulkan handles
from Vulkan memory objects such that the underlying resources can be
referenced outside the scope of the Vulkan logical device that created
them.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html):

  - [VkExternalMemoryBufferCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryBufferCreateInfoKHR.html)

- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html):

  - [VkExternalMemoryImageCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfoKHR.html)

- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html):

  - [VkExportMemoryAllocateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_EXTERNAL_MEMORY_EXTENSION_NAME`

- `VK_KHR_EXTERNAL_MEMORY_SPEC_VERSION`

- `VK_QUEUE_FAMILY_EXTERNAL_KHR`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) How do applications correlate two physical devices across process or
Vulkan instance boundaries?

**RESOLVED**: New device ID fields have been introduced by
[`VK_KHR_external_memory_capabilities`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_capabilities.html).
These fields, combined with the existing
[VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`driverVersion`
field can be used to identify compatible devices across processes,
drivers, and APIs.
[VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`pipelineCacheUUID`
is not sufficient for this purpose because despite its description in
the specification, it need only identify a unique pipeline cache format
in practice. Multiple devices may be able to use the same pipeline cache
data, and hence it would be desirable for all of them to have the same
pipeline cache UUID. However, only the same concrete physical device can
be used when sharing memory, so an actual unique device ID was
introduced. Further, the pipeline cache UUID was specific to Vulkan, but
correlation with other, non-extensible APIs is required to enable
interoperation with those APIs.

2\) If memory objects are shared between processes and APIs, is this
considered aliasing according to the rules outlined in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-memory-aliasing"
target="_blank" rel="noopener">Memory Aliasing</a> section?

**RESOLVED**: Yes. Applications must take care to obey all restrictions
imposed on aliased resources when using memory across multiple Vulkan
instances or other APIs.

3\) Are new image layouts or metadata required to specify image layouts
and layout transitions compatible with non-Vulkan APIs, or with other
instances of the same Vulkan driver?

**RESOLVED**: Separate instances of the same Vulkan driver running on
the same GPU should have identical internal layout semantics, so
applications have the tools they need to ensure views of images are
consistent between the two instances. Other APIs will fall into two
categories: Those that are Vulkan- compatible, and those that are
Vulkan-incompatible. Vulkan-incompatible APIs will require the image to
be in the GENERAL layout whenever they are accessing them.

Note this does not attempt to address cross-device transitions, nor
transitions to engines on the same device which are not visible within
the Vulkan API. Both of these are beyond the scope of this extension.

4\) Is a new barrier flag or operation of some type needed to prepare
external memory for handoff to another Vulkan instance or API and/or
receive it from another instance or API?

**RESOLVED**: Yes. Some implementations need to perform additional cache
management when transitioning memory between address spaces and other
APIs, instances, or processes which may operate in a separate address
space. Options for defining this transition include:

- A new structure that can be added to the `pNext` list in
  [VkMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryBarrier.html),
  [VkBufferMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryBarrier.html), and
  [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier.html).

- A new bit in [VkAccessFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html) that can be set to
  indicate an “external” access.

- A new bit in [VkDependencyFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlags.html)

- A new special queue family that represents an “external” queue.

A new structure has the advantage that the type of external transition
can be described in as much detail as necessary. However, there is not
currently a known need for anything beyond differentiating between
external and internal accesses, so this is likely an over-engineered
solution. The access flag bit has the advantage that it can be applied
at buffer, image, or global granularity, and semantically it maps pretty
well to the operation being described. Additionally, the API already
includes `VK_ACCESS_MEMORY_READ_BIT` and `VK_ACCESS_MEMORY_WRITE_BIT`
which appear to be intended for this purpose. However, there is no
obvious pipeline stage that would correspond to an external access, and
therefore no clear way to use `VK_ACCESS_MEMORY_READ_BIT` or
`VK_ACCESS_MEMORY_WRITE_BIT`.
[VkDependencyFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlags.html) and
[VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags.html) operate at command
granularity rather than image or buffer granularity, which would make an
entire pipeline barrier an internal→external or external→internal
barrier. This may not be a problem in practice, but seems like the wrong
scope. Another downside of [VkDependencyFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlags.html)
is that it lacks inherent directionality: there are no `src` and `dst`
variants of it in the barrier or dependency description semantics, so
two bits might need to be added to describe both internal→external and
external→internal transitions. Transitioning a resource to a special
queue family corresponds well with the operation of transitioning to a
separate Vulkan instance, in that both operations ideally include
scheduling a barrier on both sides of the transition: Both the releasing
and the acquiring queue or process. Using a special queue family
requires adding an additional reserved queue family index. Re-using
`VK_QUEUE_FAMILY_IGNORED` would have left it unclear how to transition a
concurrent usage resource from one process to another, since the
semantics would have likely been equivalent to the currently-ignored
transition of `VK_QUEUE_FAMILY_IGNORED` → `VK_QUEUE_FAMILY_IGNORED`.
Fortunately, creating a new reserved queue family index is not invasive.

Based on the above analysis, the approach of transitioning to a special
“external” queue family was chosen.

5\) Do internal driver memory arrangements and/or other internal driver
image properties need to be exported and imported when sharing images
across processes or APIs.

**RESOLVED**: Some vendors claim this is necessary on their
implementations, but it was determined that the security risks of
allowing opaque metadata to be passed from applications to the driver
were too high. Therefore, implementations which require metadata will
need to associate it with the objects represented by the external
handles, and rely on the dedicated allocation mechanism to associate the
exported and imported memory objects with a single image or buffer.

6\) Most prior interoperation and cross-process sharing APIs have been
based on image-level sharing. Should Vulkan sharing be based on
memory-object sharing or image sharing?

**RESOLVED**: These extensions have assumed memory-level sharing is the
correct granularity. Vulkan is a lower-level API than most prior APIs,
and as such attempts to closely align with to the underlying primitives
of the hardware and system-level drivers it abstracts. In general, the
resource that holds the backing store for both images and buffers of
various types is memory. Images and buffers are merely metadata
containing brief descriptions of the layout of bits within that memory.

Because memory object-based sharing is aligned with the overall Vulkan
API design, it enables the full range of Vulkan capabilities with
external objects. External memory can be used as backing for sparse
images, for example, whereas such usage would be awkward at best with a
sharing mechanism based on higher-level primitives such as images.
Further, aligning the mechanism with the API in this way provides some
hope of trivial compatibility with future API enhancements. If new
objects backed by memory objects are added to the API, they too can be
used across processes with minimal additions to the base external memory
APIs.

Earlier APIs implemented interop at a higher level, and this
necessitated entirely separate sharing APIs for images and buffers. To
co-exist and interoperate with those APIs, the Vulkan external sharing
mechanism must accommodate their model. However, if it can be agreed
that memory-based sharing is the more desirable and forward-looking
design, legacy interoperation constraints can be considered another
reason to favor memory-based sharing: while native and legacy driver
primitives that may be used to implement sharing may not be as low-level
as the API here suggests, raw memory is still the least common
denominator among the types. Image-based sharing can be cleanly derived
from a set of base memory- object sharing APIs with minimal effort,
whereas image-based sharing does not generalize well to buffer or
raw-memory sharing. Therefore, following the general Vulkan design
principle of minimalism, it is better to expose interopability with
image-based native and external primitives via the memory sharing API,
and place sufficient limits on their usage to ensure they can be used
only as backing for equivalent Vulkan images. This provides a consistent
API for applications regardless of which platform or external API they
are targeting, which makes development of multi-API and multi-platform
applications simpler.

7\) Should Vulkan define a common external handle type and provide
Vulkan functions to facilitate cross-process sharing of such handles
rather than relying on native handles to define the external objects?

**RESOLVED**: No. Cross-process sharing of resources is best left to
native platforms. There are myriad security and extensibility issues
with such a mechanism, and attempting to re-solve all those issues
within Vulkan does not align with Vulkan’s purpose as a graphics API. If
desired, such a mechanism could be built as a layer or helper library on
top of the opaque native handle defined in this family of extensions.

8\) Must implementations provide additional guarantees about state
implicitly included in memory objects for those memory objects that may
be exported?

**RESOLVED**: Implementations must ensure that sharing memory objects
does not transfer any information between the exporting and importing
instances and APIs other than that required to share the data contained
in the memory objects explicitly shared. As specific examples, data from
previously freed memory objects that used the same underlying physical
memory, and data from memory objects using adjacent physical memory must
not be visible to applications importing an exported memory object.

9\) Must implementations validate external handles the application
provides as inputs to memory import operations?

**RESOLVED**: Implementations must return an error to the application if
the provided memory handle cannot be used to complete the requested
import operation. However, implementations need not validate handles are
of the exact type specified by the application.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-10-20 (James Jones)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_external_memory"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

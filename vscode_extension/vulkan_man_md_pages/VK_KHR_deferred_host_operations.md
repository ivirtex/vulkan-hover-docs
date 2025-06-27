# VK_KHR_deferred_host_operations(3) Manual Page

## Name

VK_KHR_deferred_host_operations - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

269

## <a href="#_revision" class="anchor"></a>Revision

4

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Josh Barczak <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_deferred_host_operations%5D%20@jbarczak%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_deferred_host_operations%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jbarczak</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-11-12

**IP Status**  
No known IP claims.

**Contributors**  
- Joshua Barczak, Intel

- Jeff Bolz, NVIDIA

- Daniel Koch, NVIDIA

- Slawek Grajewski, Intel

- Tobias Hector, AMD

- Yuriy O’Donnell, Epic

- Eric Werness, NVIDIA

- Baldur Karlsson, Valve

- Jesse Barker, Unity

- Contributors to VK_KHR_acceleration_structure,
  VK_KHR_ray_tracing_pipeline

## <a href="#_description" class="anchor"></a>Description

The
[`VK_KHR_deferred_host_operations`](VK_KHR_deferred_host_operations.html)
extension defines the infrastructure and usage patterns for deferrable
commands, but does not specify any commands as deferrable. This is left
to additional dependent extensions. Commands **must** not be deferred
unless the deferral is specifically allowed by another extension which
depends on
[`VK_KHR_deferred_host_operations`](VK_KHR_deferred_host_operations.html).

## <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreateDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDeferredOperationKHR.html)

- [vkDeferredOperationJoinKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDeferredOperationJoinKHR.html)

- [vkDestroyDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDeferredOperationKHR.html)

- [vkGetDeferredOperationMaxConcurrencyKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeferredOperationMaxConcurrencyKHR.html)

- [vkGetDeferredOperationResultKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeferredOperationResultKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_DEFERRED_HOST_OPERATIONS_EXTENSION_NAME`

- `VK_KHR_DEFERRED_HOST_OPERATIONS_SPEC_VERSION`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_DEFERRED_OPERATION_KHR`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_OPERATION_DEFERRED_KHR`

  - `VK_OPERATION_NOT_DEFERRED_KHR`

  - `VK_THREAD_DONE_KHR`

  - `VK_THREAD_IDLE_KHR`

## <a href="#_code_examples" class="anchor"></a>Code Examples

The following examples will illustrate the concept of deferrable
operations using a hypothetical example. The command
`vkDoSomethingExpensive` denotes a deferrable command.

The following example illustrates how a vulkan application might request
deferral of an expensive operation:

``` c
// create a deferred operation
VkDeferredOperationKHR hOp;
VkResult result = vkCreateDeferredOperationKHR(device, pCallbacks, &hOp);
assert(result == VK_SUCCESS);

result = vkDoSomethingExpensive(device, hOp, ...);
assert( result == VK_OPERATION_DEFERRED_KHR );

// operation was deferred.  Execute it asynchronously
std::async::launch(
    [ hOp ] ( )
    {
        vkDeferredOperationJoinKHR(device, hOp);

        result = vkGetDeferredOperationResultKHR(device, hOp);

        // deferred operation is now complete.  'result' indicates success or failure

        vkDestroyDeferredOperationKHR(device, hOp, pCallbacks);
    }
);
```

The following example illustrates extracting concurrency from a single
deferred operation:

``` c
// create a deferred operation
VkDeferredOperationKHR hOp;
VkResult result = vkCreateDeferredOperationKHR(device, pCallbacks, &hOp);
assert(result == VK_SUCCESS);

result = vkDoSomethingExpensive(device, hOp, ...);
assert( result == VK_OPERATION_DEFERRED_KHR );

// Query the maximum amount of concurrency and clamp to the desired maximum
uint32_t numLaunches = std::min(vkGetDeferredOperationMaxConcurrencyKHR(device, hOp), maxThreads);

std::vector<std::future<void> > joins;

for (uint32_t i = 0; i < numLaunches; i++) {
  joins.emplace_back(std::async::launch(
    [ hOp ] ( )
    {
        vkDeferredOperationJoinKHR(device, hOp);
                // in a job system, a return of VK_THREAD_IDLE_KHR should queue another
                // job, but it is not functionally required
    }
  ));
}

for (auto &f : joins) {
  f.get();
}

result = vkGetDeferredOperationResultKHR(device, hOp);

// deferred operation is now complete.  'result' indicates success or failure

vkDestroyDeferredOperationKHR(device, hOp, pCallbacks);
```

The following example shows a subroutine which guarantees completion of
a deferred operation, in the presence of multiple worker threads, and
returns the result of the operation.

``` c
VkResult FinishDeferredOperation(VkDeferredOperationKHR hOp)
{
    // Attempt to join the operation until the implementation indicates that we should stop

    VkResult result = vkDeferredOperationJoinKHR(device, hOp);
    while( result == VK_THREAD_IDLE_KHR )
    {
        std::this_thread::yield();
        result = vkDeferredOperationJoinKHR(device, hOp);
    }

    switch( result )
    {
    case VK_SUCCESS:
        {
            // deferred operation has finished.  Query its result.
            result = vkGetDeferredOperationResultKHR(device, hOp);
        }
        break;

    case VK_THREAD_DONE_KHR:
        {
            // deferred operation is being wrapped up by another thread
            //  wait for that thread to finish
            do
            {
                std::this_thread::yield();
                result = vkGetDeferredOperationResultKHR(device, hOp);
            } while( result == VK_NOT_READY );
        }
        break;

    default:
        assert(false); // other conditions are illegal.
        break;
    }

    return result;
}
```

## <a href="#_issues" class="anchor"></a>Issues

1.  Should this extension have a VkPhysicalDevice\*FeaturesKHR
    structure?

**RESOLVED**: No. This extension does not add any functionality on its
own and requires a dependent extension to actually enable functionality
and thus there is no value in adding a feature structure. If necessary,
any dependent extension could add a feature boolean if it wanted to
indicate that it is adding optional deferral support.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-12-05 (Josh Barczak, Daniel Koch)

  - Initial draft.

- Revision 2, 2020-03-06 (Daniel Koch, Tobias Hector)

  - Add missing VK_OBJECT_TYPE_DEFERRED_OPERATION_KHR enum

  - fix sample code

  - Clarified deferred operation parameter lifetimes (#2018,!3647)

- Revision 3, 2020-05-15 (Josh Barczak)

  - Clarify behavior of vkGetDeferredOperationMaxConcurrencyKHR,
    allowing it to return 0 if the operation is complete (#2036,!3850)

- Revision 4, 2020-11-12 (Tobias Hector, Daniel Koch)

  - Remove VkDeferredOperationInfoKHR and change return value semantics
    when deferred host operations are in use (#2067,3813)

  - clarify return value of vkGetDeferredOperationResultKHR
    (#2339,!4110)

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_deferred_host_operations"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
